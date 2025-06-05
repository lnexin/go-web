package gee

import (
	"fmt"
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(m string, pattern string, handler HandlerFunc) {
	log.Printf("Route Add %4s - %s", m, pattern)
	key := m + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if h, ok := r.handlers[key]; ok {
		h(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}
