package main

import (
	"net/http"

	"github.com/minicloudsky/cloudsky/cloudsky"
)

func main() {
	r := cloudsky.Default()
	r.GET("/", func(c *cloudsky.Context) {
		c.String(http.StatusOK, "Hello cloudskyktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *cloudsky.Context) {
		names := []string{"cloudskyktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
