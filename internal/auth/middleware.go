package auth

import "net/http"

type Middleware struct{}

func NewMiddleware() Middleware {
	return Middleware{}
}

func (m Middleware) RequireAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
