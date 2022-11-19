package main

import (
	"github.com/gin-gonic/gin"
)

// АРГУМЕНТЫ ВСЕХ КОНТАКТОВ РАССЫЛОК, КОТОРЫЕ СОДЕРЖАТ ДАННЫЕ
var client = []Client{}

// АРГУМЕНТЫ ВСЕХ ОСНОВНЫХ КОНТАКТОВ РАССЫЛОК, КОТОРЫЕ СОДЕРЖАТ ДАННЫЕ
var main_contact = []Main_contact{}

// АРГУМЕНТЫ ДОСТУПОВ HAPPYFOX
var staff_access = []Staff_access{}

// АРГУМЕНТЫ ОТЧЁТА HAPPYFOX
var report = []Report{}

// СОЗДАЁМ МЕТОДЫ ЗАПРОСОВ ПО API
func main() {
	router := gin.Default()
	router.GET("/api/clients", getClients)
	router.GET("/api/main_contacts", getClientMain)
	router.GET("/api/staff_access", getStaffAccess)
	router.GET("/api/report", getReport)
	router.POST("/api/add_client", postClients)
	router.POST("/api/add_main_contact", postClientMain)
	router.POST("/api/add_staff_access", postStaffAccess)
	router.POST("/api/add_report", postReport)
	router.PUT("/api/add_client", putClients)

	// УКАЗЫВАЕМ АДРЕС ДЛЯ СЕРВИСА
	router.Run("0.0.0.0:8055")
}
