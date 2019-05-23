package core

import (
	"github.com/gobuffalo/packr"
)

//Templates wraps packr and handled embedding the /static files in the executable
var Templates = packr.NewBox("../resources/templates")
