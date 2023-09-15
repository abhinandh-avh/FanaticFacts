package routes

import (
	"new-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", controllers.GetHomePage)

	r.GET("/service", controllers.GetServicePage)

	r.GET("/contact", controllers.GetContactPage)

	r.GET("/about", controllers.GetAboutPage)

	r.GET("/signup", controllers.GetSignPage)

	r.GET("/login", controllers.GetLoginPage)

	r.POST("/login", controllers.PostLogin)
	r.POST("/signup", controllers.PostSignin)

	// ...

	return r

}
