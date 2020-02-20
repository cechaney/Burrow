package controllers

import (
	"fmt"
	"net/http"

	"github.com/cechaney/burrow/pkg/core"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "OK")

}

//GetHealthController builds and returns the healthController
func GetHealthController() core.Controller {

	controller := core.Controller{
		Name:    "health",
		Path:    "/health",
		Handler: healthHandler,
	}

	return controller

}
