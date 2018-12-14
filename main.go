package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	zipname := os.Getenv("TESTDATA_PATH")

	// todo
	log.Println("create http server")
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("starting server ...")
	r.Run()
}
