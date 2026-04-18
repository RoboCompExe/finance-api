package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	id := c.Param("id")

	var balance int64
	err := DB.QueryRow("SELECT balance FROM accounts WHERE id=$1", id).Scan(&balance)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
