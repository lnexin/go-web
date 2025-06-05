package main

import (
	"main/gee"
	"net/http"
)

func main() {
	engine := gee.New()
	//engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprint(w, "")
	//})
	//
	//engine.POST("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//
	//	m := make(map[string]string)
	//
	//	for s, strings := range request.Header {
	//		m[s] = fmt.Sprintf("%q", strings)
	//		//fmt.Fprintf(writer, "Header[%q] = %s\n", s, strings)
	//	}
	//	fmt.Fprintf(writer, "Header = %v\n", m)
	//})

	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>")
	})

	engine.GET("/hello", func(context *gee.Context) {
		context.String(http.StatusOK, "hello %s, you uri is %s\n", context.Query("name"), context.Path)
	})

	engine.POST("/login", func(context *gee.Context) {
		context.JSON(200, gee.HM{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	engine.Run(":8080")
}
