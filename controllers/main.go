package controllers

import "github.com/cechaney/burrow/core"

//GetControllers returns the full array of active controllers
func GetControllers() []core.Controller {

	/*
		Add all of the controllers you want active in the app to this array
	*/
	controllers := []core.Controller{
		GetHTMLController(),
		GetJSONController(),
		//TODO: add Mustache controller
	}

	return controllers
}
