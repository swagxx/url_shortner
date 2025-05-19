package middleware

import "net/http"

type WriteResponse struct {
	http.ResponseWriter
	StatusCode  int
	wroteHeader bool
}

func (w *WriteResponse) WriteHeader(code int) {
	if !w.wroteHeader {
		w.StatusCode = code
		w.ResponseWriter.WriteHeader(code)
		w.wroteHeader = true
	}
}
