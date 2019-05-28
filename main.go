package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/cechaney/burrow/controllers"
	"github.com/cechaney/burrow/core"
	"github.com/cechaney/burrow/middleware"
)

func main() {

	core.ConfigureLogger(core.Config)
	core.ConfigureStatic(core.Static)
	core.AttachRouter(core.Config.GetString("context"))
	core.ConfigureMiddleware(middleware.GetMiddleware())
	core.ConfigureControllers(controllers.GetControllers())

	serverConfig, err := getServerConfig()

	if err != nil {
		core.Logger.Error("app startup failed")
		core.Logger.Error(err)
		panic(err)
	}

	//Configure and start the app
	app := &http.Server{
		Handler:      core.Router,
		Addr:         serverConfig.Address + ":" + serverConfig.Port,
		WriteTimeout: serverConfig.WriteTimeout,
		ReadTimeout:  serverConfig.ReadTimeout,
	}

	core.Logger.Infof("App started on port %s", serverConfig.Port)

	core.Logger.Fatal(app.ListenAndServe())

}

func getServerConfig() (core.ServerConfig, error) {

	var config core.ServerConfig

	//Get the configured port
	port := core.Config.GetString("port")

	if len(port) <= 0 {
		return config, errors.New("no server port specified")
	}

	address := core.Config.GetString("address")

	if len(address) <= 0 {
		return config, errors.New("no server address specified")
	}

	writeTimeout, err := time.ParseDuration(
		core.Config.GetString("writeTimeout"))

	if err != nil {
		return config, errors.New("no server write timeout specified")
	}

	readTimeout, _ := time.ParseDuration(
		core.Config.GetString("readTimeout"))

	if err != nil {
		return config, errors.New("no server read timeout specified")
	}

	config.Port = port
	config.Address = address
	config.WriteTimeout = writeTimeout
	config.ReadTimeout = readTimeout

	return config, nil

}
