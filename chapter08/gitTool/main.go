package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter08/gitTool/model"
	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
	"github.com/urfave/cli"
)

var GITHUB_TOKEN string
var requestOptions *grequests.RequestOptions

const (
	GITHUB_API_URL = "https://api.github.com"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	GITHUB_TOKEN = os.Getenv("GITHUB_PERSONAL_TOKEN")
	requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func createGist(url string, args []string) *grequests.Response {
	description := args[0]
	var fileContents = make(map[string]model.File)
	for i := 1; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filenames. Absolute path (or) same directory are allowed")
			return nil
		}
		var file model.File
		file.Content = string(data)
		fileContents[filepath.Base(args[i])] = file
	}
	var gist = model.Gist{Description: description, Public: true, Files: fileContents}
	log.Println(gist)
	var postBody, _ = json.Marshal(gist)
	var requestOptions_copy = requestOptions
	requestOptions_copy.JSON = string(postBody)
	resp, err := grequests.Post(url, requestOptions_copy)
	if err != nil {
		log.Println("Create request failed for Github API")
	}
	return resp
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "fetch",
			Aliases: []string{"f"},
			Usage: "Fetch the repo details with user. [Usage]: githubAPI fetch user",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					var repos []model.Repo
					user := c.Args()[0]
					var repoURL = fmt.Sprintf("%s/users/%s/repos", GITHUB_API_URL, user)
					resp := getStats(repoURL)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Please give a username. See -h to see help")
				}
				return nil
			},
		},
		{
			Name: "create",
			Aliases: []string{"c"},
			Usage: "Creates a gist from the given text. [Usage]: githubAPI name 'description' sample.txt",
			Action: func (c * cli.Context) error {
				if c.NArg() > 1 {
					args := c.Args()
					var postURL = fmt.Sprintf("%s/gists", GITHUB_API_URL)
					resp := createGist(postURL, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}

	app.Version = "1.0"
	app.Run(os.Args)
}