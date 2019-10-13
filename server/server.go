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
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

const defaultPort = "8080"

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "DB_USER",
			EnvVar: "DB_USER",
			Value:  "postgres",
			Usage:  "DB user",
		},
		cli.StringFlag{
			Name:   "DB_PASSWORD",
			EnvVar: "DB_PASSWORD",
			Value:  "postgres",
			Usage:  "DB password",
		},
		cli.StringFlag{
			Name:   "DB_HOST",
			EnvVar: "DB_HOST",
			Value:  "localhost",
			Usage:  "DB url",
		},
		cli.StringFlag{
			Name:   "DB_PORT",
			EnvVar: "DB_PORT",
			Value:  "5432",
			Usage:  "DB url",
		},
		cli.StringFlag{
			Name:   "DATABASE_NAME",
			EnvVar: "DATABASE_NAME",
			Value:  "test",
			Usage:  "DB url",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		port := defaultPort

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		})

		db := store.NewStore(&pg.Options{
			User:     ctx.String("DB_USER"),
			Password: ctx.String("DB_PASSWORD"),
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%s", ctx.String("DB_HOST"), ctx.String("DB_PORT")),
			Database: ctx.String("DATABASE_NAME"),
		})

		fs := http.FileServer(http.Dir("out"))
		http.Handle("/", http.StripPrefix("/", fs))

		http.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
		http.Handle("/query", c.Handler(handler.GraphQL(bara.NewExecutableSchema(bara.Config{Resolvers: &bara.Resolver{DB: db}}),
			handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
				debug.PrintStack()
				return errors.New("user message on panic")
			}),
		)))

		log.Printf("connect to http://localhost:%s/payground for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
