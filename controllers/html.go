package controllers

import (
	"fmt"
	"net/http"

	"github.com/cechaney/burrow/core"
)

//htmlController returns a handler that outputs HTML
func htmlHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html><body><h1>This is HTML!</h1></body></html>")
}

//GetHTMLController builds and returns the HelloHandler
func GetHTMLController() core.Controller {

	controller := core.Controller{
		Name:    "html",
		Path:    "/html",
		Handler: htmlHandler,
	}

	return controller

}
