package main

import (
	"github.com/minicloudsky/cloudsky/cloudsky"
	"net/http"
)

func main() {
	r := cloudsky.New()
	r.GET("/", func(c *cloudsky.Context) {
		c.HTML(http.StatusOK, "<h1>Hello cloudsky</h1>")
	})
	r.GET("/hello", func(c *cloudsky.Context) {
		// expect /hello?name=cloudskyktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *cloudsky.Context) {
		c.JSON(http.StatusOK, cloudsky.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
