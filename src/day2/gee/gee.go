package gee

import (
	"net/http"
)

type Engine struct {
	router *router
}

type HandlerFunc func(*Context)

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
	context := newContext(w, req)
	engine.router.handle(context)

	//key := req.Method + "-" + req.URL.Path
	//fmt.Println(engine.router)
	//fmt.Println(key)
	//if handlerFunc, ok := engine.router[key]; ok {
	//	handlerFunc(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 NOT Ex2ist URI: %q\n", req.URL.Path)
	//}
}
