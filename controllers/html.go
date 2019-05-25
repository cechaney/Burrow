package controllers

import (
	"net/http"
	"html/template"

	"github.com/cechaney/burrow/core"
)

type pagedata struct{
	Message string
}

//htmlController returns a handler that outputs HTML
func htmlHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	data := pagedata{"Hi!"}

	content := `
	<html>
		<head>
			<link rel='shortcut icon' type='image/x-icon' href='/favicon.ico' />
		</head>
		<body>
			<h1>{{.Message}}</h1>
		</body>
	</html>
	`
	t, err := template.New("html").Parse(content)

	if err != nil {
		
		core.Logger.Error(err)

		http.Error(w, "", http.StatusInternalServerError)

	} else {
		t.Execute(w, data)
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
