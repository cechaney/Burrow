package main

import (
	"net/http"

	"github.com/cechaney/burrow/core"
)

func main() {

	core.Logger.Info("Booting up...")

	http.Handle("/", http.FileServer(core.Static))

	port := core.Config.GetString("port")

	core.Logger.Info("Listening on port: " + port)

	http.ListenAndServe(":"+port, nil)

}
