package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"log"
	"os"
)

var (
	Acc map[int]Account
)

func main() {
	Acc = make(map[int]Account)

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

	r.POST("/accounts/new/", CreateUser)
	r.GET("/accounts/filter/", FilterUser)

	log.Println("starting server ...")
	_ = r.Run(os.Getenv("SERVER_PORT"))
}

func CreateUser(c *gin.Context) {
	var json Account
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	Acc[json.ID] = json

	c.JSON(201, gin.H{})
}

func FilterUser(c *gin.Context) {

	r := funk.Filter(Acc, func(x Account) bool {

	})

}
