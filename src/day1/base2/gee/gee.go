package gee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]HandlerFunc
}

type HandlerFunc func(
	http.ResponseWriter,
	*http.Request,
)

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	fmt.Println(engine.router)
	fmt.Println(key)
	if handlerFunc, ok := engine.router[key]; ok {
		handlerFunc(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT Ex2ist URI: %q\n", req.URL.Path)
	}
}
