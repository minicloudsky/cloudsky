package cloudsky

import (
	"net/http"
	"testing"
)

func TestCloudsky(t *testing.T) {
	c := Default()
	c.GET("/ping", func(context *Context) {
		context.JSON(http.StatusOK, "pong")
	})
	err := c.Run(":9000")
	if err != nil {
		return
	}
}
