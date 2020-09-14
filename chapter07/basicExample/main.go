package main

import (
	"log"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter07/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
