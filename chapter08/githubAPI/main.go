package main

import (
	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
	"log"
	"os"
)

var GITHUB_TOKEN string
var requestOptions *grequests.RequestOptions

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Forks int `json:"forks"`
	Private bool `json:"private"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	GITHUB_TOKEN = os.Getenv("GITHUB_PERSONAL_TOKEN")
	requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}
}

func main() {
	var repos []Repo
	var repoURL = "http://api.github.com/users/d4n13l-4lf4/repos"
	resp := getStats(repoURL)
	resp.JSON(&repos)
	log.Println(repos)
}

