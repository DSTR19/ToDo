package core_http_response

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

var (
	StatusCodeUnitialized = -1
)

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUnitialized,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriter) GetStatusCodeOrPanic() int {
	if w.statusCode == StatusCodeUnitialized {
		panic("status code not initialized")
	}
	return w.statusCode
}
