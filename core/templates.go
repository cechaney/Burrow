package core

import (
	"github.com/gobuffalo/packr"
)

//Templates wraps packr and handled embedding the /static files in the executable
var Templates = packr.NewBox("../resources/templates")

//We should probably try to pre-compile all of the templates
//https://golang.org/doc/articles/wiki/#tmp_10
