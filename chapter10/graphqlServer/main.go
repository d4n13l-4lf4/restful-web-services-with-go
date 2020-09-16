package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

type Player struct {
	ID int `json:"int"`
	Name string `json:"name"`
	HighScore int `json:"highScore"`
	IsOnline bool `json:"isOnline"`
	Location string `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

var players = []Player{
	{ID: 123, Name: "Daniel", HighScore: 1100, IsOnline: true, Location: "Somewhere"},
	{ID: 124, Name: "Alfredo", HighScore: 1101, IsOnline: true, Location: "Somewhere"},
}

var playerObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"highScore": &graphql.Field{
			Type: graphql.String,
		},
		"isOnline": &graphql.Field{
			Type: graphql.Boolean,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"levelsUnlocked": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

func main() {
	fields := graphql.Fields{
		"players": &graphql.Field{
			Type: graphql.NewList(playerObject),
			Description: "All players",
			Resolve: func(p graphql.ResolveParams) (interface{}, error){
				return players, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", nil)
}