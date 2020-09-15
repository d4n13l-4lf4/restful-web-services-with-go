package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your name")
var age = flag.Int("age", 0, "your age")

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years), Welcome to the command line world", *name, *age)
}