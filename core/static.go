package core

import (
	"github.com/gobuffalo/packr"
)

//Static wraps packr and handled embedding the /static files in the executable
var Static = initStatic()

func initStatic() packr.Box {

	box := packr.NewBox("../resources/static")

	return box

}
