package web_transport_http

import (
	core_http_server "github.com/PhitonBedrosovich/golang-todoapp/internal/core/transport/http/server"
)

type WebHTTPHandler struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
}

func NewWebHTTPHandler(
	webService WebService,
) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/", // достаточно будет написать домен сайта, лиюо ip с портом сервера и в ответ будет отдана html-страница
			Handler: h.GetMainPage,
		},
	}
}
