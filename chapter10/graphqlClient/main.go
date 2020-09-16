package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"log"
	"os"
)

type Response struct {
	License struct{
		Name string `json:"name"`
		Description string `json:"description"`
	}	`json:"license"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	githubToken := os.Getenv("GITHUB_PERSONAL_TOKEN")
	graphqlEndpoint := os.Getenv("GITHUB_GRAPHQL_ENDPOINT")
	client := graphql.NewClient(graphqlEndpoint)
	req := graphql.NewRequest(`
			query {
				license(key: "apache-2.0") {
					name
					description
				}	
			}
		`)
	req.Header.Add("Authorization", "bearer " + githubToken)
	ctx := context.Background()

	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData.License.Description)
}