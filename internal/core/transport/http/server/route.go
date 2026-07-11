package core_http_server

import "net/http"

// Route - набор параметров, благодаря которым мультиплексор сможет понять
// как по входящим данным http-запроса выбрать для этого http-запроса http-обработчик
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewRoute(
	method  string,
	path    string,
	handler http.HandlerFunc,
) Route {
	return Route {
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
