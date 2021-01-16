package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "JWT_SECRET",
			EnvVar: "JWT_SECRET",
			Value:  "secret",
			Usage:  "JWT Secret for the api's authentification.",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":  1,
			"role": "admin",
		})

		tokenString, err := token.SignedString([]byte(ctx.String("JWT_SECRET")))

		fmt.Println(tokenString, err)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
