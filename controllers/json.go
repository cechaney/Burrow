package controllers

import (
	"encoding/json"
	"net/http"
	
	"github.com/cechaney/burrow/core"
)

//jsonController returns a handler that outputs HTML
func jsonHandler(w http.ResponseWriter, r *http.Request) {

	type Message struct {
		Content string `json:"content"`
	}

	message := Message{Content: "This is JSON!"}

	output, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		w.Header().Set("Content-Type", "application/json")

		w.Write(output)
	}

}

//GetJSONController builds and returns the HelloHandler
func GetJSONController() core.Controller {

	controller := core.Controller{
		Name:    "json",
		Path:    "/json",
		Handler: jsonHandler,
	}

	return controller

}
