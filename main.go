package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static assets
	r.Static("/assets", "./assets")

	// Define routes and handlers
	r.GET("/", func(c *gin.Context) {
		renderTemplate(c, "index.html", nil)
	})
	r.GET("/services-content", func(c *gin.Context) {
		renderTemplate(c, "service.html", nil)
	})

	// ...

	r.Run(":8080")
}

// Function to render HTML templates
func renderTemplate(c *gin.Context, tmplName string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + tmplName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load template"})
		return
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		return
	}
}
