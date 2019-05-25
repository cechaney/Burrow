package controllers

import (
	"net/http"

	"github.com/cechaney/burrow/core"
)

//faviconHandler returns a handler that returns the favicon for the app
func faviconHandler(w http.ResponseWriter, r *http.Request) {

	favicon, err := core.Static.Find("favicon.ico")

	if err != nil{
		http.Error(w, "", http.StatusNotFound)
	} else {

		w.Header().Set("Content-Type", "image/x-icon")
		
		w.Write(favicon)
	}
	
}

//GetFaviconController builds and returns the faviconHandler
func GetFaviconController() core.Controller {

	controller := core.Controller{
		Name:    "favicon",
		Path:    "/favicon.ico",
		Handler: faviconHandler,
	}

	return controller

}
