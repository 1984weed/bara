package main

import (
	"bara"
	"bara/auth"
	"bara/generated"
	"bara/problem/executor"
	problem_repository "bara/problem/repository"
	problem_resolver "bara/problem/resolver"
	problem_usecase "bara/problem/usecase"
	user_endpoint "bara/user/endpoint"
	user_repository "bara/user/repository"
	user_resolver "bara/user/resolver"
	user_usecase "bara/user/usecase"

	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
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

		codeExecutor := executor.NewExecutorClient(ctx.Bool("WITHOUT_CONTAINER"), 10)

		// Problem
		problemRepo := problem_repository.NewProblemRepositoryRunner(db)
		problemUc := problem_usecase.NewProblemUsecase(problemRepo, codeExecutor, timeoutContext)
		problemResolver := problem_resolver.NewProblemResolver(problemUc)

		// User
		userRepoRunner := user_repository.NewUserRepositoryRunner(db)
		userUc := user_usecase.NewUserUsecase(userRepoRunner, timeoutContext)
		userResolver := user_resolver.NewUserResolver(userUc)
		userEndpoint := user_endpoint.NewUserEndpoint(userUc, store)

		router := chi.NewRouter()
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		router.Use(cors.Handler)
		router.Use(auth.Middleware(userRepoRunner, store))

		router.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
		router.Handle("/query", cors.Handler(handler.GraphQL(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &bara.Resolver{
					DB:              db,
					ProblemResolver: problemResolver,
					UserResolver:    userResolver,
				},
			}),
			handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
				debug.PrintStack()
				return errors.New("An error happens")
			}),
		)))

		router.HandleFunc("/signup", userEndpoint.SignUp)
		router.HandleFunc("/login", userEndpoint.Login)
		router.HandleFunc("/logout", userEndpoint.Logout)

		log.Printf("connect to http://localhost:%s/payground for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
