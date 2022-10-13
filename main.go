package main

import (
	_ "assignment-2/docs"
	"assignment-2/router"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	r := router.StartApp()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
