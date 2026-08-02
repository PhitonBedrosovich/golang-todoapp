package main

import (
	"crypto/subtle"
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:embed public/login.html
var loginFS embed.FS

const cookieName = "todoapp_token"

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	OK       bool   `json:"ok"`
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type tokenClaims struct {
	Username string `json:"sub"`
	jwt.RegisteredClaims
}

type server struct {
	username     string
	password     string
	jwtSecret    []byte
	tokenTTL     time.Duration
	cookieSecure bool
	upstream     *url.URL
	proxy        http.Handler
	loginHTML    []byte
}

func main() {
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	secret := os.Getenv("JWT_SECRET")
	upstreamRaw := envOr("UPSTREAM", "http://todoapp:5050")
	addr := envOr("HTTP_ADDR", ":8080")

	if username == "" || password == "" || secret == "" {
		log.Fatal("AUTH_USERNAME, AUTH_PASSWORD and JWT_SECRET are required")
	}

	upstream, err := url.Parse(upstreamRaw)
	if err != nil {
		log.Fatalf("invalid UPSTREAM: %v", err)
	}

	loginHTML, err := fs.ReadFile(loginFS, "public/login.html")
	if err != nil {
		log.Fatalf("read login.html: %v", err)
	}

	s := &server{
		username:     username,
		password:     password,
		jwtSecret:    []byte(secret),
		tokenTTL:     parseDuration(envOr("JWT_EXPIRY", "168h")),
		cookieSecure: strings.EqualFold(os.Getenv("COOKIE_SECURE"), "true"),
		upstream:     upstream,
		loginHTML:    loginHTML,
	}
	s.proxy = s.newProxy()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /login", s.handleLoginPage)
	mux.HandleFunc("POST /auth/login", s.handleLogin)
	mux.HandleFunc("POST /auth/logout", s.handleLogout)
	mux.HandleFunc("GET /auth/me", s.handleMe)
	mux.HandleFunc("/", s.handleProxy)

	log.Printf("auth-gateway listening on %s, upstream %s", addr, upstreamRaw)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseDuration(raw string) time.Duration {
	d, err := time.ParseDuration(raw)
	if err != nil {
		return 168 * time.Hour
	}
	return d
}

func (s *server) newProxy() http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(s.upstream)
	origDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		origDirector(req)
		req.Host = s.upstream.Host
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		http.Error(w, "upstream unavailable", http.StatusBadGateway)
	}
	return proxy
}

func (s *server) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	if s.tokenFromRequest(r) != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s.loginHTML)
}

func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, loginResponse{Message: "invalid request body"})
		return
	}

	if !s.validCredentials(req.Username, req.Password) {
		writeJSON(w, http.StatusUnauthorized, loginResponse{Message: "неверный логин или пароль"})
		return
	}

	token, err := s.issueToken(req.Username)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, loginResponse{Message: "failed to create token"})
		return
	}

	s.setAuthCookie(w, token)
	writeJSON(w, http.StatusOK, loginResponse{OK: true, Username: req.Username})
}

func (s *server) handleLogout(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
	writeJSON(w, http.StatusOK, loginResponse{OK: true})
}

func (s *server) handleMe(w http.ResponseWriter, r *http.Request) {
	claims, ok := s.parseToken(s.tokenFromRequest(r))
	if !ok {
		writeJSON(w, http.StatusUnauthorized, loginResponse{Message: "unauthorized"})
		return
	}
	writeJSON(w, http.StatusOK, loginResponse{OK: true, Username: claims.Username})
}

func (s *server) handleProxy(w http.ResponseWriter, r *http.Request) {
	if s.tokenFromRequest(r) == "" {
		if isAPIRequest(r) {
			writeJSON(w, http.StatusUnauthorized, loginResponse{Message: "unauthorized"})
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if _, ok := s.parseToken(s.tokenFromRequest(r)); !ok {
		s.clearAuthCookie(w)
		if isAPIRequest(r) {
			writeJSON(w, http.StatusUnauthorized, loginResponse{Message: "unauthorized"})
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	s.proxy.ServeHTTP(w, r)
}

func (s *server) validCredentials(username, password string) bool {
	userOK := subtle.ConstantTimeCompare([]byte(username), []byte(s.username)) == 1
	passOK := subtle.ConstantTimeCompare([]byte(password), []byte(s.password)) == 1
	return userOK && passOK
}

func (s *server) issueToken(username string) (string, error) {
	now := time.Now()
	claims := tokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.tokenTTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.jwtSecret)
}

func (s *server) parseToken(token string) (tokenClaims, bool) {
	if token == "" {
		return tokenClaims{}, false
	}
	parsed, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return s.jwtSecret, nil
	})
	if err != nil || !parsed.Valid {
		return tokenClaims{}, false
	}
	claims, ok := parsed.Claims.(*tokenClaims)
	if !ok {
		return tokenClaims{}, false
	}
	return *claims, true
}

func (s *server) tokenFromRequest(r *http.Request) string {
	if auth := r.Header.Get("Authorization"); strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (s *server) setAuthCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   int(s.tokenTTL.Seconds()),
		HttpOnly: true,
		Secure:   s.cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

func (s *server) clearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.cookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

func isAPIRequest(r *http.Request) bool {
	return strings.HasPrefix(r.URL.Path, "/api/")
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
