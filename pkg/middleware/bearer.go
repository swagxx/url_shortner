package middleware

import (
	"context"
	"judo/configs"
	"judo/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "emailKey"
)

func writeHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	http.Error(w, "unauthorized", http.StatusUnauthorized)
	return
}

func Bearer(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if !strings.HasPrefix(bearer, "Bearer ") {
			log.Println("bearer token required")
			writeHeader(w)
			return
		}

		if bearer == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(bearer, "Bearer ")
		if token == "" {
			next.ServeHTTP(w, r)
			return

		}
		data, ok := jwt.NewJWT(config.Auth.Secret).ParseToken(token)

		if !ok || data == nil {
			writeHeader(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
