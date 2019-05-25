package core

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/handlers"
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
func ConfigureMiddleware(middlewares []mux.MiddlewareFunc) {

	for _, m := range middlewares {
		Router.Use(m)
	}

}

//AttachRouter attached the handler to a specified path
func AttachRouter(path string) {
	http.Handle(path, Router)
}

//ConfigureControllers registers an array of controllers to the router
func ConfigureControllers(controllers []Controller) {

	for _, c := range controllers {

		compressed := handlers.CompressHandler(c.Handler)

		Router.Handle(c.Path, compressed)
	}

}
