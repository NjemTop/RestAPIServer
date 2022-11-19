package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ОТПРАВКА ТЕЛО ЗАПРОСА PUT В ФОРМАТЕ JSON
func putClients(c *gin.Context) {
	var updateClient Client

	if err := c.BindJSON(&updateClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is bad"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "200",
		"data":   updateClient,
	})
}
