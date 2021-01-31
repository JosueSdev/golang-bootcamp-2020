package router

import "net/http"

//jsonContenTypeMiddleware returns a middleware that adds content type json by default to responses
func jsonContenTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
