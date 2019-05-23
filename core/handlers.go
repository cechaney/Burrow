package core

import "net/http"

//BurrowHandler wraps an http.Handler with extra details about name and path
type BurrowHandler struct {
	Name    string
	Path    string
	Handler http.HandlerFunc
}
