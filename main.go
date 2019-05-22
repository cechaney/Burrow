package main

import (
	"net/http"

	"github.com/cechaney/burrow/core"
)

func main() {

	core.Logger.Info("We are go!")

	http.Handle("/", http.FileServer(core.Static))

	port := core.Config.GetString("port")

	if len(port) <= 0 {
		port = core.Config.GetString("ENV_PORT")
	}

	core.Logger.Info("Listening on port: " + port)

	http.ListenAndServe(":"+port, nil)

}
