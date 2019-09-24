package main

import (
	"bara"
	"bara/store"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := store.NewStore(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Network:  "tcp",
		Addr:     "localhost:5432",
		Database: "bara",
	})

	fmt.Println(db)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(bara.NewExecutableSchema(bara.Config{Resolvers: &bara.Resolver{DB: db}}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
