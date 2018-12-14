package main

import (
	"github.com/deissh/HighLoad18/handler"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := LoadFromZip(os.Getenv("ZIP_FILE"))

	if err != nil {
		log.Panic("fault when unzip file")
	}

	log.Println("create http server")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g := r.Group("/accounts")
	{
		g.GET("/filter", handler.FilterHandler)
	}

	log.Println("starting server ...")
	_ = r.Run()
}
