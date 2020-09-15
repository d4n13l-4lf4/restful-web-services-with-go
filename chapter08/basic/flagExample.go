package main

import (
	"flag"
	"log"
)

var name string

func init() {
	flag.StringVar(&name, "name", "stranger", "your name")
}

func main() {
	flag.Parse()
	log.Printf("Hello %s, Welcome to the command line world", name)
}