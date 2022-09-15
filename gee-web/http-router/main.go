package main

import (
	"example/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s,you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
