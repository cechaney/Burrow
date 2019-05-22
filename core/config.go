package core

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Config wraps viper so we don't have snakes all over our app
var Config = initConfig()

/*
	Configuration supports 4 different patterns
	===========================================================================
	1. Environment variables: Just ENV_PORT for now.  Add more as needed
	2. JSON config file: File location set with the -config cmd line option
	3. External config service: Not implemented yet
*/
func initConfig() *viper.Viper {

	//Get the port set by environment variable
	viper.BindEnv("ENV_PORT")

	//Get config file from command line arguments.  Default is "local.json"
	configArg := flag.String("config", "local", "Config file location")

	flag.Parse()

	viper.SetConfigName(*configArg)

	viper.AddConfigPath("./resources/config/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	//Watch the config file and update the config after a change
	viper.WatchConfig()

	//Let's log when the config updates
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})

	return viper.GetViper()

}
