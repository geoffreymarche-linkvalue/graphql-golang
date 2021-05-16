package main

import (
	"graphql-golang/app"
	"graphql-golang/mutations"
	"graphql-golang/queries"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	config := app.Load()
	dbService := app.NewDBService(config)
	if dbService != nil {

	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/", httpHandler)
	log.Print("ready: listening...\n")

	err = http.ListenAndServe(":8383", nil)
	if err != nil {
		panic(err)
	}
}