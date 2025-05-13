package middleware

import "net/http"

type WriteResponse struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WriteResponse) WriteHeader(code int) {
	w.StatusCode = code
}
