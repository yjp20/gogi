package gogi

import (
	"github.com/julienschmidt/httprouter"
)

type AuthMethodHandler struct {
	Method  string
	Slug    string
	Handler httprouter.Handle
}

type AuthMethod interface {
	Client() (string, string)
	Handlers(*Context) []AuthMethodHandler
}
