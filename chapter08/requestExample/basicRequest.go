package main

import (
	"github.com/levigross/grequests"
	"log"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(resp.String())
}