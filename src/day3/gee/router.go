package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := ParsePattern(pattern)

	key := fmt.Sprintf("%s-%s", method, pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	searhParts := ParsePattern(path)

	params := make(map[string]string)
	targetRoot, ok := r.roots[method]
	if !ok {
		return nil, nil

	}
	targetNode := targetRoot.search(searhParts, 0)
	if targetNode == nil {
		return nil, nil
	}
	// != nil

	targetParts := ParsePattern(targetNode.pattern)
	for i, part := range targetParts {
		if part[0] == ':' {
			params[part[1:len(part)]] = searhParts[i]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searhParts[i:], "/")
			break
		}

	}
	return targetNode, params
}

func (r *Router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)

	if node != nil {
		c.params = params
		key := c.Method + "-" + node.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}

//func newTestRouter() *Router {
//	router := newRouter()
//	router.addRoute("GET", "/", nil)
//	router.addRoute("GET", "/hello", nil)
//	router.addRoute("GET", "/hello/:name", nil)
//	router.addRoute("GET", "/hello/b/c", nil)
//	router.addRoute("GET", "/hi/:name", nil)
//	router.addRoute("GET", "/assert/*filepath", nil)
//	return router
//}git a
//
//func TestParsePatter(t *testing.T) {
//	ok := reflect.DeepEqual(ParsePattern("/p/:name"), []string{"p", ":name"})
//	ok = reflect.DeepEqual(ParsePattern("/p/*"), []string{"p", "*"})
//	ok = reflect.DeepEqual(ParsePattern("/p/*name/*"), []string{"p", "*name"})
//	if !ok {
//		t.Fatal("tests ParsePattern field")
//	}
//}
//
//func TestGetRoute(t *testing.T) {
//	router := newTestRouter()
//	n, ps := router.getRoute("GET", "/hello/geeeee")
//	if n == nil{
//		t.Fatal("nil should not be return")
//	}
//
//	if n.pattern != "/hello/:name"{
//		t.Fatal("should match /hello/:name")
//	}
//
//	if ps["name"] != "geeeee" {
//		t.Fatal("name should not be equal to 'geeeee'")
//	}
//	fmt.Println("match path: %s, params[name]: %s\n", n.pattern, ps["name"])
//
//}
