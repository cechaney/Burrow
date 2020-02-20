package controllers

import (
	"html/template"
	"net/http"

	"github.com/cechaney/burrow/pkg/core"
)

type htmlPageData struct {
	Message string
}

//htmlController returns a handler that outputs HTML
func htmlHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	data := htmlPageData{"Have fun digging around fellow gophers."}

	content, err := core.Templates.FindString("index.html")

	if err != nil {

		core.Logger.Error(err)

		http.Error(w, "", http.StatusInternalServerError)

	} else {

		t, err := template.New("html").Parse(content)

		if err != nil {

			core.Logger.Error(err)

			http.Error(w, "", http.StatusInternalServerError)

		} else {
			t.Execute(w, data)
		}

	}

}

//GetHTMLController builds and returns the htmlHandler
func GetHTMLController() core.Controller {

	controller := core.Controller{
		Name:    "html",
		Path:    "/html",
		Handler: htmlHandler,
	}

	return controller

}
