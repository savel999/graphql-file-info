package main

import (
	"graphql-file-info/graph"
	"graphql-file-info/graph/generated"
	"graphql-file-info/graph/repository"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newResolver() *graph.Resolver {
	itemsRep := repository.NewListItemsRepository()
	itemsRep.InitDefaultList()

	return &graph.Resolver{ItemsRepository: itemsRep}
}
