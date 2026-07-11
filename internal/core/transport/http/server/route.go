package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/PhitonBedrosovich/golang-todoapp/internal/core/transport/http/middleware"
)

// Route - набор параметров, благодаря которым мультиплексор сможет понять
// как по входящим данным http-запроса выбрать для этого http-запроса http-обработчик
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return  core_http_middleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}
