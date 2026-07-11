package core_http_response

import "net/http"

var (
	StatusCodeUninitialized = -1
)

// оборачивает стандартный ResponseWriter и дополнительно запоминать статус код
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode: 	StatusCodeUninitialized,
	}
}

// вызываем метод с фактической записью статус кода ответа и запоминаем статус код
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

// метод для получения значений statusCode
func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == StatusCodeUninitialized {
		panic("no status code set")
	}
	
	return rw.statusCode
}