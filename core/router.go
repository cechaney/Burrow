package core

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

//Router wraps the Gorilla mux so we don't have apes all over our app
var Router = mux.NewRouter()

func init() {

}

//ConfigureStatic takes a packr.Box to use for /static/ path assets
func ConfigureStatic(static packr.Box) {

	//Set up serving of static assets
	Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(static)))

}

//ConfigureMiddleware registers an array of middleware to the router
func ConfigureMiddleware() {

}

//ConfigureHandlers registers an array of handlers to router
func ConfigureHandlers() {

}

//AttachRouter attached the handler to a specified path
func AttachRouter(path string) {
	http.Handle(path, Router)
}
