package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()

	r := gin.Default()

	r.GET("/balance/:id", GetBalance)
	r.POST("/transfer", Transfer)

	r.Run(":8080")
}
