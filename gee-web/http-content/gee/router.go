package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s -%s", method, pattern)
	k := method + "-" + pattern
	r.handlers[k] = handler
}

func (r *router) handle(c *Context) {
	k := c.Method + "-" + c.Path
	if handler, ok := r.handlers[k]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
