package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/cechaney/burrow/core"
)

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		core.Logger.Info(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

//GetAccessLogMiddleware returnd the xssMiddleware
func GetAccessLogMiddleware() mux.MiddlewareFunc{
	return accessLogMiddleware
}
