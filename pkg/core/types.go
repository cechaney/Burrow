package core

import (
	"net/http"
	"time"
)

//ServerConfig represents http Server configuration properties
type ServerConfig struct {
	Port         string
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

//Controller wraps an http.Handler with extra details about name and path
type Controller struct {
	Name    string
	Path    string
	Handler http.HandlerFunc
}
