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

	var frontendOrigin string
	if env.AppEnv == "DEV" {
		frontendOrigin = "http://127.0.0.1:5173"
	} else if env.AppEnv == "PROD" {
		frontendOrigin = "https://ocwcentral.com"
	} else {
		panic("invalid APP_ENV")
	}

	resolver := InitializeResolver()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		srv.ServeHTTP(w, r)
	})
	// Handle playground for DEV environment
	if env.AppEnv == "DEV" {
		http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/playground for GraphQL playground", env.Port)
	}
	log.Fatal(http.ListenAndServe(":"+env.Port, nil))
}
