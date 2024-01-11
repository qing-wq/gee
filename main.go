package main

import (
	"fmt"
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", indexHandler)
	r.POST("/hello", helloHandler)

	r.Run(":8080")
}

// handler echoes r.URL.Path
func indexHandler(c *gee.Context) {
	fmt.Println(c.Path)
}

// handler echoes r.URL.Header
func helloHandler(c *gee.Context) {
	for k, v := range c.Req.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}
}
