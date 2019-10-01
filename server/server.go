package main

import (
	"bara"
	"bara/store"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowCredentials: true,
	})

	db := store.NewStore(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Network:  "tcp",
		Addr:     "localhost:5432",
		Database: "bara",
	})

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(handler.GraphQL(bara.NewExecutableSchema(bara.Config{Resolvers: &bara.Resolver{DB: db}}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
