package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/PhitonBedrosovich/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/PhitonBedrosovich/golang-todoapp/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

// структура HTTPServer, HTTPServer будет строится поверх стандартного мультиплексора из пакета http
type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger

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

func (h *HTTPServer) RequesterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

// метод будет запускать HTTPServer, обработчик http запросов и возвращать ошибки
func (h *HTTPServer) Run(ctx context.Context) error {
	// навесим на мультиплексор middlewares
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	// канал ошибок
	ch := make(chan error, 1)

	// горутина
	go func() {
		defer close(ch)

		h.log.Warn("start HTTP server", zap.String("addr", h.config.Addr))

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
		h.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			// ждем пока все http-обработчики доработают до конца
			h.config.ShutdownTimeout,
		)
		defer cancel()

		// просим HTTP server завершится аккуратно, то есть сервер больше не будет принимать новые http-запросы
		// если запущенные http-обработчики не успевают доработать за ShutdownTimeout, тогда у нас метод Shutdown вернет ошибку
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")
	}

	return nil
}
