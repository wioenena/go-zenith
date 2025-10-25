package controller

import "github.com/wioenena/go-zenith/http"

type handlerCache map[http.HttpMethod]map[string][]HandlerFunc

type BaseController struct {
	handlerCache handlerCache
}

func (c *BaseController) Get(path string, handlers ...HandlerFunc) {
	c.ByMethod(path, http.GET, handlers...)
}

func (c *BaseController) Post(path string, handlers ...HandlerFunc) {
	c.ByMethod(path, http.POST, handlers...)
}

func (c *BaseController) Put(path string, handlers ...HandlerFunc) {
	c.ByMethod(path, http.PUT, handlers...)
}

func (c *BaseController) Patch(path string, handlers ...HandlerFunc) {
	c.ByMethod(path, http.PATCH, handlers...)
}

func (c *BaseController) Delete(path string, handlers ...HandlerFunc) {
	c.ByMethod(path, http.DELETE, handlers...)
}

func (c *BaseController) ByMethod(path string, method http.HttpMethod, handlers ...HandlerFunc) {
	if _, exists := c.handlerCache[method][path]; !exists {
		c.handlerCache[method][path] = make([]HandlerFunc, 0)
	}
	c.handlerCache[method][path] = append(c.handlerCache[method][path], handlers...)
}

func (c *BaseController) initializeHandlerCache() {
	c.handlerCache = make(handlerCache)
	c.handlerCache[http.GET] = make(map[string][]HandlerFunc)
	c.handlerCache[http.POST] = make(map[string][]HandlerFunc)
	c.handlerCache[http.PUT] = make(map[string][]HandlerFunc)
	c.handlerCache[http.PATCH] = make(map[string][]HandlerFunc)
	c.handlerCache[http.DELETE] = make(map[string][]HandlerFunc)
}

func (c *BaseController) GetHandlers(method http.HttpMethod, path string) ([]HandlerFunc, error) {
	if handlers, exists := c.handlerCache[method][path]; exists {
		return handlers, nil
	}

	return nil, ErrControllerHandlersEmpty
}

func NewBaseController() *BaseController {
	c := new(BaseController)
	c.initializeHandlerCache()
	return c
}
