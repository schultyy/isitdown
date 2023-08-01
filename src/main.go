package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from %v", "Gin")
	})

	r.POST("/perform_check", func(c *gin.Context) {
		url := c.PostForm("website_url")
		fmt.Printf("url: %s\n", url)
		stdout, stderr, err := Shellout(fmt.Sprintf("ping -c 4 %s", url))
		commandOutput := stdout + stderr
		fmt.Printf("OUTPUT: %s", commandOutput)
		fmt.Printf("OUTPUT ERR: %s", err)
		c.HTML(http.StatusOK, "results.html", gin.H{
			"answer": commandOutput,
			"error":  err,
		})
	})
	r.Run(":3000")
}
