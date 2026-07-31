package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PhitonBedrosovich/golang-todoapp/docs"
	core_logger "github.com/PhitonBedrosovich/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/PhitonBedrosovich/golang-todoapp/internal/core/transport/http/middleware"
	"go.uber.org/zap"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// структура HTTPServer, HTTPServer будет строится поверх стандартного мультиплексора из пакета http
type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    *core_logger.Logger

	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router.WithMiddleware()),
		)
	}
}

// регистрация роутев напрямую в HTTP сервере в обход APIVersionRouter
func (s *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		s.mux.Handle(pattern, route.WithMiddleware())
	}
}

func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
			httpSwagger.DefaultModelsExpandDepth(-1),
		),
	)

	s.mux.HandleFunc(
		"/swagger/doc.json",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
		},
	)
}

// метод будет запускать HTTPServer, обработчик http запросов и возвращать ошибки
func (s *HTTPServer) Run(ctx context.Context) error {
	// навесим на мультиплексор middlewares
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middleware...)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	// канал ошибок
	ch := make(chan error, 1)

	// горутина
	go func() {
		defer close(ch)

		s.log.Warn("start HTTP server", zap.String("addr", s.config.Addr))

		// обрабатываем http-запросы, которые летят на сервер
		err := server.ListenAndServe()

		// если это не ошибка о том, что сервер корректно завершился, тогда запишем в канал ошибок ....
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	// прочитаем из канала ошибку и тогда весь метод Run вернет ошибку, например: listen and server HTTP
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			// ждем пока все http-обработчики доработают до конца
			s.config.ShutdownTimeout,
		)
		defer cancel()

		// просим HTTP server завершится аккуратно, то есть сервер больше не будет принимать новые http-запросы
		// если запущенные http-обработчики не успевают доработать за ShutdownTimeout, тогда у нас метод Shutdown вернет ошибку
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.log.Warn("HTTP server stopped")
	}

	return nil
}
