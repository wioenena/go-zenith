package controller

import (
	"github.com/wioenena/go-zenith/http"
)

type HandlerFunc func(ctx http.HttpContext) error

type Controller interface {
	Setup()
	GetRoute() string
	Get(path string, handlers ...HandlerFunc)
	Post(path string, handlers ...HandlerFunc)
	Put(path string, handlers ...HandlerFunc)
	Patch(path string, handlers ...HandlerFunc)
	Delete(path string, handlers ...HandlerFunc)
	ByMethod(path string, method http.HttpMethod, handlers ...HandlerFunc)
	GetHandlers(method http.HttpMethod, path string) ([]HandlerFunc, error)
}
