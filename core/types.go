package core

import(
	"net/http"
)

//Controller wraps an http.Handler with extra details about name and path
type Controller struct {
	Name    string
	Path    string
	Handler http.HandlerFunc
}