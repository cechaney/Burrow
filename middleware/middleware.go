package middleware

import (
	"github.com/gorilla/mux"
)

//GetMiddleware returns the full array of active controllers
func GetMiddleware() []mux.MiddlewareFunc{

	/*
		Add all of the middleware you want active in the app to this array
	*/
	mws := []mux.MiddlewareFunc{
		GetAccessLogMiddleware(),
	}

	return mws
}