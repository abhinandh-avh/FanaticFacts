package main

import (
	"new-project/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run(":8080")
}
