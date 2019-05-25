package main

import (
	"net/http"
	"time"

	"github.com/cechaney/burrow/middleware"
	"github.com/cechaney/burrow/controllers"
	"github.com/cechaney/burrow/core"
)

func main() {

	core.ConfigureLogger(core.Config)
	core.ConfigureStatic(core.Static)
	core.AttachRouter(core.Config.GetString("context"))
	core.ConfigureMiddleware(middleware.GetMiddleware())
	core.ConfigureControllers(controllers.GetControllers())

	//Get the configured port
	port := core.Config.GetString("port")

	if len(port) <= 0 {
		core.Logger.Panic("No port specified")
		panic("Failed to start because no port configured")
	}

	//Configure and start the app
	app := &http.Server{
		Handler:      core.Router,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	core.Logger.Infof("App started on port %s", port)

	core.Logger.Fatal(app.ListenAndServe())

}

func startApp() {

}
