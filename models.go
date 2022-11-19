package main

// СОЗДАЁМ АРГУМЕНТЫ ДЛЯ JSON НА СЕРВИСЕ
type Client struct {
	Client string `json:"Client"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
}

//
type Main_contact struct {
	Client      string `json:"Client"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	MainContact string `json:"Main_contact"`
}

type Staff_access struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Active  bool   `json:"active"`
	StaffID int    `json:"Staff_ID"`
}

type Report struct{}
