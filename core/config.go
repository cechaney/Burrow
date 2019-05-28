package core

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/spf13/viper"
)

const portEnvVarName = "ENV_BURROW_PORT"

//Config wraps viper so we don't have snakes all over our app
var Config *viper.Viper

func init() {

	/*
		Configuration supports 3 different patterns
		===========================================================================
		1. Environment variables: Just ENV_PORT for now.  Add more as needed
		2. JSON config file: File location set with the -config cmd line option
		3. External config service: Not implemented yet
	*/

	//Standardize on JSON config files
	viper.SetConfigType("json")

	//Get config file from command line arguments.  Default is "local.json"
	configArg := flag.String("config", "local", "Config file location")

	flag.Parse()

	configBox := packr.NewBox("../resources/config")

	configBytes, findConfigError := configBox.Find(*configArg + ".json")

	if findConfigError != nil {
		panic(fmt.Errorf("%s", findConfigError))
	}

	configReader := bytes.NewReader(configBytes)

	readConfigError := viper.ReadConfig(configReader)

	if readConfigError != nil {
		panic(fmt.Errorf("%s", readConfigError))
	}

	//Get the port set by environment variable
	viper.BindEnv(portEnvVarName)

	envPort := viper.GetString(portEnvVarName)

	if len(envPort) > 0 {
		viper.Set("port", envPort)
	}

	Config = viper.GetViper()

}
