package main

import (
	"log"

	"HTTP-REST-API/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); != nil {
		log.Fatal(err)
	}
}
