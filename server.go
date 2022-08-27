package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kafugen/ocwcentral/env"
	"github.com/kafugen/ocwcentral/graph/generated"
)

func main() {
	env := env.NewEnvConfig()
	resolver := InitializeResolver()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Port)
	log.Fatal(http.ListenAndServe(":"+env.Port, nil))
}
