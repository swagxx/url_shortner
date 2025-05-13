package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Bearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(bearer, "Bearer ")
		if token == "" {
			next.ServeHTTP(w, r)
			return

		}
		next.ServeHTTP(w, r)
		fmt.Println(token)
	})
}
