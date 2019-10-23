package main

import (
	"bara"
	"bara/generated"
	"bara/problem/repository"
	"bara/problem/resolver"
	"bara/problem/usecase"

	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

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
			Value:  "bara",
			Usage:  "DB url",
		},
		cli.StringFlag{
			Name:   "PORT",
			EnvVar: "PORT",
			Value:  "8080",
			Usage:  "Web app port",
		},
		cli.BoolFlag{
			Name:   "WITHOUT_CONTAINER",
			EnvVar: "WITHOUT_CONTAINER",
			Usage:  "Can define the application is running on local",
		},
	}
	app.Action = func(ctx *cli.Context) error {

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
		})

		db := pg.Connect(
			&pg.Options{
				User:     ctx.String("DB_USER"),
				Password: ctx.String("DB_PASSWORD"),
				Network:  "tcp",
				Addr:     fmt.Sprintf("%s:%s", ctx.String("DB_HOST"), ctx.String("DB_PORT")),
				Database: ctx.String("DATABASE_NAME"),
			})

		port := ctx.String("PORT")
		timeoutContext := time.Duration(40) * time.Second

		fs := http.FileServer(http.Dir("out"))
		problemRepo := repository.NewProblemRepository(db)
		problemUc := usecase.NewProblemUsecase(problemRepo, timeoutContext)
		problemResolver := resolver.NewProblemResolver(problemUc)

		http.Handle("/", http.StripPrefix("/", fs))

		http.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
		http.Handle("/query", c.Handler(handler.GraphQL(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &bara.Resolver{
					DB:               db,
					ProblemResolver:  problemResolver,
					WithoutContainer: ctx.Bool("WITHOUT_CONTAINER"),
				},
			}),
			handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
				debug.PrintStack()
				return errors.New("An error happens")
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
