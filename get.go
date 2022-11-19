package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ОТВЕТ НА ЗАПРОС GET ПО КЛИЕНТАМ РАССЫЛКИ
func getClients(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, client)
}

// ОТВЕТ НА ЗАПРОС GET ПО ОСНОВНЫМ КЛИЕНТАМ РАССЫЛКИ
func getClientMain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, main_contact)
}

// ОТВЕТ НА ЗАПРОС GET ПО ПРАВАМ HAPPYFOX
func getStaffAccess(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, staff_access)
}

// ОТВЕТ НА ЗАПРОС GET НА ПОЛУЧЕНИЕ ОТЧЁТОВ
func getReport(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, report)
}

// ОТВЕТЫ GET НА ЗАПРОСЫ ПО ID
func getClientByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range client {
		if a.Client == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Contact not found"})
}
