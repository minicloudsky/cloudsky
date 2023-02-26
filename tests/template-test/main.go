package main

import (
	"fmt"
	"github.com/minicloudsky/cloudsky/cloudsky"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := cloudsky.New()
	r.Use(cloudsky.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "cloudskyktutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *cloudsky.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *cloudsky.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", cloudsky.H{
			"title":  "cloudsky",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *cloudsky.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", cloudsky.H{
			"title": "cloudsky",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
