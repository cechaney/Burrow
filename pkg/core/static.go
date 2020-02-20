package core

import (
	"github.com/gobuffalo/packr"
)

//Static wraps packr and handled embedding the /static files in the executable
var Static = packr.NewBox("../resources/static")
