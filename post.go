package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ОТПРАВКА ТЕЛО ЗАПРОСА POST В ФОРМАТЕ JSON
func postClients(c *gin.Context) {
	var newClient Client

	if err := c.BindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is bad"})
		return
	}

	client = append(client, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}

// ОТПРАВКА ТЕЛО ЗАПРОСА POST В ФОРМАТЕ JSON
func postClientMain(c *gin.Context) {
	var newMain_contacts Main_contact

	if err := c.BindJSON(&newMain_contacts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is bad"})
		return
	}

	main_contact = append(main_contact, newMain_contacts)
	c.IndentedJSON(http.StatusCreated, newMain_contacts)
}

// ОТПРАВКА ТЕЛО ЗАПРОСА POST В ФОРМАТЕ JSON
func postStaffAccess(c *gin.Context) {
	var newStaff_access Staff_access

	if err := c.BindJSON(&newStaff_access); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is bad"})
		return
	}

	staff_access = append(staff_access, newStaff_access)
	c.IndentedJSON(http.StatusCreated, newStaff_access)
}

// ОТПРАВКА ТЕЛО ЗАПРОСА POST В ФОРМАТЕ JSON
func postReport(c *gin.Context) {
	var newReport Report

	if err := c.BindJSON(&newReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is bad"})
		return
	}

	report = append(report, newReport)
	c.IndentedJSON(http.StatusCreated, newReport)
}
