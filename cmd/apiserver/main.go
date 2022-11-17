package main

func main() {
	s := apiserver.New()
	if err := s.Start(); != nil {
		log.Fatal(err)
	}
}
