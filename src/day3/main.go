package main

import (
	"main/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello!!!!</h1>")
	})

	r.GET("/hello", func(context *gee.Context) {
		context.String(200, "hello %s, you are at %s\n", context.Query("name"), context.Path)
	})

	r.GET("/hello/:name", func(context *gee.Context) {

		context.String(200, "hello %s, you param[name] is %s, path: %s\n", context.Param("name"), context.Param("name"), context.Path)

	})

	r.GET("/assert/*filepath", func(context *gee.Context) {
		context.JSON(200, gee.HM{"filepath": context.Param("filepath")})
	})
	r.Run(":8080")

}
