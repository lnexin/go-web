package main

import (
	"fmt"
	"main/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %s\n", request.URL.Path)
	})

	r.POST("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for s, strings := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %s\n", s, strings)
		}
	})

	r.Run(":8080")
}
