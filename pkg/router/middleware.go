package router

import (
	"context"
	"net/http"
)

// mwAddContext adds a context to each request
func mwAddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "debugging", "itsdebuggingtime")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
