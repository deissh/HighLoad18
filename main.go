package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	Acc []Account
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := LoadFromZip(os.Getenv("ZIP_FILE"))

	if err != nil {
		log.Panic("fault when unzip file")
	}

	sort.Slice(Acc, func(i, j int) bool {
		return Acc[i].ID > Acc[j].ID
	})

	log.Println("create http server")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g := r.Group("/accounts")
	{
		g.POST("/:name/", PostHandler)
		g.GET("/:param", GetHandler)
		g.GET("/:param/", GetHandler)
		g.GET("/:param/:action/", GetHandler)
		//g.GET("/filter/", FilterUser)
		//g.GET("/group", Group)
		//g.GET("/*id/recommend/", Recommend)
		//g.GET("/:id/suggest/", Suggest)
	}

	log.Println("starting server ...")
	_ = r.Run(os.Getenv("SERVER_PORT"))
}

func PostHandler(c *gin.Context) {
	param := c.Param("name")
	if param == "new" {
		CreateUser(c)
		return
	}
	if param == "likes" {
		Liks(c)
		return
	}

	UpdateUser(c)
}

func GetHandler(c *gin.Context) {
	param := c.Param("param")
	action := c.Param("action")
	if param == "filter" {
		FilterUser(c)
		return
	}
	if param == "group" {
		Group(c)
		return
	}
	if action == "suggest" {
		Suggest(c)
		return
	}
	if action == "recommend" {
		Recommend(c)
		return
	}

	c.Status(404)
}

func CreateUser(c *gin.Context) {
	var json Account
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	//Acc = append(Acc, json)

	c.JSON(201, gin.H{})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("name"))
	if err != nil {
		c.JSON(404, gin.H{})
		return
	}

	r := funk.Find(Acc, func(item Account) bool {
		return item.ID == id
	})

	if r == nil {
		c.JSON(404, gin.H{})
		return
	}

	var json Account
	err = c.BindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(202, gin.H{})
}

func Liks(c *gin.Context) {
	var json Likes
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(202, gin.H{})
}

func Group(c *gin.Context) {
	c.JSON(200, gin.H{})
}
