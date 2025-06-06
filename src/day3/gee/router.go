package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Roots    map[string]*Node       `json:"Roots"`
	Handlers map[string]HandlerFunc `json:"Handlers"`
}

func newRouter() *Router {
	return &Router{
		Roots:    make(map[string]*Node),
		Handlers: make(map[string]HandlerFunc),
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

// PrintTree 递归打印树形结构
func PrintTree(r *Router) {
	//marshal, err := json.MarshalIndent(r, "", "  ")
	marshal, err := json.MarshalIndent(r.Roots, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	PrintTree(r)
	parts := ParsePattern(pattern)

	key := fmt.Sprintf("%s-%s", method, pattern)
	_, ok := r.Roots[method]
	if !ok {
		r.Roots[method] = &Node{}
	}
	fmt.Println("----------------------------")
	fmt.Println("method: ", method, "pattern: ", pattern)
	fmt.Println("----------------------------")
	r.Roots[method].insert(pattern, parts, 0)
	r.Handlers[key] = handler
	PrintTree(r)
	fmt.Println("====================================")
	//log.Println("====================================")
}

func (r *Router) getRoute(method string, path string) (*Node, map[string]string) {
	searhParts := ParsePattern(path)

	params := make(map[string]string)
	targetRoot, ok := r.Roots[method]
	if !ok {
		return nil, nil

	}
	targetNode := targetRoot.search(searhParts, 0)
	if targetNode == nil {
		return nil, nil
	}
	// != nil

	targetParts := ParsePattern(targetNode.Pattern)
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
		key := c.Method + "-" + node.Pattern
		r.Handlers[key](c)
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
//	if n.Pattern != "/hello/:name"{
//		t.Fatal("should match /hello/:name")
//	}
//
//	if ps["name"] != "geeeee" {
//		t.Fatal("name should not be equal to 'geeeee'")
//	}
//	fmt.Println("match path: %s, params[name]: %s\n", n.Pattern, ps["name"])
//
//}
