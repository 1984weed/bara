package main

import (
	"bara"
	"bara/auth"
	contest_repository "bara/contest/repository"
	contest_resolver "bara/contest/resolver"
	contest_usecase "bara/contest/usecase"
	"bara/generated"
	"bara/problem/executor"
	problem_repository "bara/problem/repository"
	problem_resolver "bara/problem/resolver"
	problem_rest "bara/problem/rest"
	problem_usecase "bara/problem/usecase"
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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-chi/chi"
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
			Name:   "DB_PASS",
			EnvVar: "DB_PASS",
			Value:  "postgres",
			Usage:  "DB password",
		},
		cli.StringFlag{
			Name:   "DB_SOCKET_PATH",
			EnvVar: "DB_SOCKET_PATH",
			Value:  "/cloudsql",
			Usage:  "DB_SOCKET_PATH for cloud run",
		},
		cli.StringFlag{
			Name:   "INSTANCE_CONNECTION_NAME",
			EnvVar: "INSTANCE_CONNECTION_NAME",
			Value:  "",
			Usage:  "INSTANCE_CONNECTION_NAME for cloud run",
		},
		cli.StringFlag{
			Name:   "DB_HOST",
			EnvVar: "DB_HOST",
			Value:  "",
			Usage:  "DB url",
		},
		cli.StringFlag{
			Name:   "DB_PORT",
			EnvVar: "DB_PORT",
			Value:  "5432",
			Usage:  "DB url",
		},
		cli.StringFlag{
			Name:   "DB_NAME",
			EnvVar: "DB_NAME",
			Value:  "bara",
			Usage:  "DB name",
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
		cli.StringFlag{
			Name:   "REDIS_URL",
			EnvVar: "REDIS_URL",
			Value:  "redis://127.0.0.1:6379",
			Usage:  "Redis url",
		},
		cli.StringFlag{
			Name:   "FRONT_URL",
			EnvVar: "FRONT_URL",
			Value:  "http://localhost:3000",
			Usage:  "Front url",
		},
		cli.StringFlag{
			Name:   "AWS_S3_ACCESS_TOKEN",
			EnvVar: "AWS_S3_ACCESS_TOKEN",
			Value:  "",
			Usage:  "AWS access token",
		},
		cli.StringFlag{
			Name:   "AWS_S3_SECRET_TOKEN",
			EnvVar: "AWS_S3_SECRET_TOKEN",
			Value:  "",
			Usage:  "AWS secret token",
		},
		cli.StringFlag{
			Name:   "S3_REGION",
			EnvVar: "S3_REGION",
			Value:  "test",
			Usage:  "S3 region",
		},
		cli.StringFlag{
			Name:   "AWS_S3_BUCKET_NAME",
			EnvVar: "AWS_S3_BUCKET_NAME",
			Value:  "test",
			Usage:  "S3 bucket name",
		},
		cli.StringFlag{
			Name:   "AWS_S3_ACCOUNT_FOLDER",
			EnvVar: "AWS_S3_ACCOUNT_FOLDER",
			Value:  "test",
			Usage:  "S3 folder",
		},
		cli.StringFlag{
			Name:   "GCP_BUCKET",
			EnvVar: "GCP_BUCKET",
			Value:  "",
			Usage:  "GCP bucket",
		},
		cli.StringFlag{
			Name:   "GCP_PROJECTID",
			EnvVar: "GCP_PROJECTID",
			Value:  "",
			Usage:  "GCP project id",
		},
		cli.StringFlag{
			Name:   "JWT_SECRET",
			EnvVar: "JWT_SECRET",
			Value:  "secret",
			Usage:  "JWT Secret for the api's authentification.",
		},
	}
	app.Action = func(ctx *cli.Context) error {

		db := pg.Connect(
			&pg.Options{
				User:     ctx.String("DB_USER"),
				Password: ctx.String("DB_PASS"),
				Network:  "tcp",
				Addr:     fmt.Sprintf("%s:%s", ctx.String("DB_HOST"), ctx.String("DB_PORT")),
				Database: ctx.String("DB_NAME"),
			})

		port := ctx.String("PORT")
		timeoutContext := time.Duration(5) * time.Second

		codeExecutor := executor.NewExecutorClient(ctx.Bool("WITHOUT_CONTAINER"), 10)

		// Contest
		contestRepo := contest_repository.NewContestRepositoryRunner(db)
		contestUc := contest_usecase.NewContestUsecase(contestRepo)
		contestResolver := contest_resolver.NewContestResolver(contestUc)

		// Problem
		problemRepo := problem_repository.NewProblemRepositoryRunner(db)
		problemUc := problem_usecase.NewProblemUsecase(problemRepo, codeExecutor, timeoutContext)
		problemResolver := problem_resolver.NewProblemResolver(problemUc, contestUc)

		// User
		userRepoRunner := user_repository.NewUserRepositoryRunner(db)

		creds := credentials.NewStaticCredentials(ctx.String("AWS_S3_ACCESS_TOKEN"), ctx.String("AWS_S3_SECRET_TOKEN"), ctx.String(""))

		s, _ := session.NewSession(&aws.Config{
			Credentials: creds,
			Region:      aws.String(ctx.String("S3_REGION")),
		})
		// User Image
		userImage := user_repository.NewUserImageRepository(s, ctx.String("GCP_BUCKET"), ctx.String("GCP_PROJECTID"))

		userUc := user_usecase.NewUserUsecase(userRepoRunner, userImage, timeoutContext)
		userResolver := user_resolver.NewUserResolver(userUc)

		router := chi.NewRouter()
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{ctx.String("FRONT_URL")},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		router.Use(cors.Handler)
		router.Use(auth.Middleware(ctx.String("JWT_SECRET")))

		router.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
		router.Handle("/query", cors.Handler(handler.GraphQL(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &bara.Resolver{
					DB:              db,
					ProblemResolver: problemResolver,
					UserResolver:    userResolver,
					ContestResolver: contestResolver,
				},
			}),
			handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
				debug.PrintStack()
				return errors.New("Unknown error happens")
			}),
		)))

		// Provide rest api to use some features
		router.Route("/v1", func(r chi.Router) {
			r.Route("/problems", func(pr chi.Router) {
				// Create a problem
				pr.Get("/", problem_rest.Helloworld)
			})
		})

		log.Printf("connect to http://localhost:%s/payground for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, router))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
