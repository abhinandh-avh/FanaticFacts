package controllers

import (
	"database/sql"
	"log"
	"new-project/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var (
	loginMessage string
	loginSuccess bool
	user         models.Userinfo
	person       models.Newuser
	db           *sql.DB
)

func GetHomePage(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func GetServicePage(c *gin.Context) {
	c.HTML(200, "service.tmpl", nil)
}

func GetContactPage(c *gin.Context) {
	c.HTML(200, "contact.tmpl", nil)
}

func GetAboutPage(c *gin.Context) {
	c.HTML(200, "about.tmpl", nil)
}

func GetSignPage(c *gin.Context) {
	c.HTML(200, "signup.tmpl", nil)
}

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{
		"Message": loginMessage,
		"Success": loginSuccess,
		"ID":      user.ID,
		"Name":    user.Name,
		"Age":     user.Age,
	})
	loginMessage = ""
	loginSuccess = false
}

func PostLogin(c *gin.Context) {
	db, err := sql.Open("mysql", "root:Abhinandh@123@tcp(localhost:3306)/campaign")
	if err != nil {
		log.Fatal("hello")
	}
	defer db.Close()
	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("hi")
	}

	user1 := c.PostForm("username")
	pass := c.PostForm("password")
	query := "SELECT id,name,age FROM users WHERE username = ? AND password = ?"

	rows, err := db.Query(query, user1, pass)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	loginSuccess = false
	loginMessage = "Login failed. Please check your credentials."
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err == nil {
			// Authentication successful
			loginSuccess = true
			loginMessage = "Login successful"
		}
	}

	c.Redirect(303, "/login")
}

func PostSignin(c *gin.Context) {
	db, err := sql.Open("mysql", "root:Abhinandh@123@tcp(localhost:3306)/campaign")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Parse form data
	if err := c.ShouldBind(&person); err != nil {
		c.HTML(400, "signup.tmpl", gin.H{"Error": err.Error()})
		return
	}

	// Insert user data into the database
	stmt, err := db.Prepare("INSERT INTO users (name, age, username, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.Exec(person.Name, person.Age, person.Username, person.Password)
	if err != nil {
		log.Fatal(err)
	}
	loginMessage = "Account Created Successfully!"
	c.Redirect(303, "/login")
}
