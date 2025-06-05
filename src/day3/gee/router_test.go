package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *Router {
	router := newRouter()
	router.addRoute("GET", "/", nil)
	router.addRoute("GET", "/hello", nil)
	router.addRoute("GET", "/hello/:name", nil)
	router.addRoute("GET", "/hello/b/c", nil)
	router.addRoute("GET", "/hi/:name", nil)
	router.addRoute("GET", "/assert/*filepath", nil)
	return router
}

func TestParsePatter(t *testing.T) {
	ok := reflect.DeepEqual(ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = reflect.DeepEqual(ParsePattern("/p/*"), []string{"p", "*"})
	ok = reflect.DeepEqual(ParsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("tests ParsePattern field")
	}
}

func TestGetRoute(t *testing.T) {
	router := newTestRouter()
	n, ps := router.getRoute("GET", "/hello/geeeee")
	if n == nil {
		t.Fatal("nil should not be return")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geeeee" {
		t.Fatal("name should not be equal to 'geeeee'")
	}
	fmt.Println("match path: %s, params[name]: %s\n", n.pattern, ps["name"])

}
