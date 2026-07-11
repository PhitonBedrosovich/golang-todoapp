package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_logger "github.com/PhitonBedrosovich/golang-todoapp/internal/core/logger"
	core_http_response "github.com/PhitonBedrosovich/golang-todoapp/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

// 1.
func RequestID() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

// 2.
func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requstID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requstID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// 3.
func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			// вызов метода PanicResponse структуры responseHandler при возникновении паники, отлавливает паники
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			// функция Panic вызывает следующий (next) handler, который она оборачивает
			next.ServeHTTP(w, r)
		})
	}
}

// 5. middleware логирует все входящие входящие запросы и все выходящие ответы
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.String("http_method", r.Method),
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r) // выполнение HTTP handlera (запроса)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Now().Sub(before)), // получаем текущее время и вычитаем время до выполнения запроса
			)
		})
	}
}
