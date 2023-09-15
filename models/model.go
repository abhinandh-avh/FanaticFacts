package models

type Userinfo struct {
	ID   int
	Name string
	Age  int
}

type Newuser struct {
	Name     string `form:"name"`
	Age      int    `form:"age"`
	Username string `form:"username"`
	Password string `form:"password"`
}
