package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.tmpl", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.tmpl", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tmpl", nil)
	})

	// ...

	r.Run(":8080")
}

