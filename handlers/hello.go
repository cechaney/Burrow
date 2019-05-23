package handlers

import (
	"fmt"
	"html"
	"net/http"

	"github.com/cechaney/burrow/core"
)

//HelloHandler says "Hello"
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getHelloHandler() core.BurrowHandler {

	helloHandler := core.BurrowHandler{
		Name:    "hello",
		Path:    "/hello",
		Handler: hello,
	}

	return helloHandler

}
