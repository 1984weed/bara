package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// SubmitCode has code users typed
type SubmitCode struct {
	TypedCode string `json:"typedCode"`
	Lang      string `json:"lang"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/:id/interpret_solution", func(c *gin.Context) {
		var submitCode SubmitCode

		if err := c.BindJSON(&submitCode); err != nil {
			return
		}

		inputCommand := fmt.Sprintf(`echo "%s" > ./temp && echo "%s" >> ./temp  && node temp`, submitCode.TypedCode, "helloworld()")
		out, err := exec.Command("docker", "run", "node:12.10.0-alpine", "/bin/ash", "-c", inputCommand).Output()

		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"result": string(out),
		})
	})
	r.Run()
}
