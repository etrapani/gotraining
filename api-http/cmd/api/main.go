package main

import (
	"example.com/gotraining/go-hexagonal_http_api-course/cmd/api/bootstrap"
	"log"


)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
