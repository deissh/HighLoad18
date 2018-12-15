package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"log"
	"os"
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

	log.Println("create http server")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/accounts/new/", CreateUser)
	r.POST("/accounts/:id/")
	r.POST("/accounts/likes/")

	r.GET("/accounts/filter/", FilterUser)
	r.GET("/accounts/group", Group)
	r.GET("/accounts/<id>/recommend", Recommend)
	r.GET("/accounts/<id>/suggest/", Suggest)

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

	Acc = append(Acc, json)

	c.JSON(201, gin.H{})
}

func FilterUser(c *gin.Context) {
	var filter Filters
	err := c.Bind(&filter)
	if err != nil {
		c.Status(400)
		return
	}

	i := 0
	limit, err := strconv.Atoi(c.DefaultQuery("limit", ""))

	if err != nil {
		c.JSON(200, gin.H{
			"accounts": Acc,
		})
		return
	}

	result := funk.Filter(Acc, func(ac Account) bool {
		f := false
		if i >= limit {
			return false
		} else {
			i++
		}

		if filter.Sex_eq == ac.Sex {
			f = true
		} else {
			if filter.Sex_eq != "" {
				f = false
			}
		}

		if filter.Status_eq == ac.Status {
			f = true
		} else {
			if filter.Status_eq != "" {
				f = false
			}
		}
		if filter.Status_neq != ac.Status {
			f = true
		} else {
			if filter.Status_neq != "" {
				f = false
			}
		}

		if filter.Fname_eq == ac.Fname {
			f = true
		} else {
			if filter.Fname_eq != "" {
				f = false
			}
		}
		if filter.Fname_any == "1" && ac.Fname == "" {
			f = true
		} else {
			if filter.Fname_any != "" {
				f = false
			}
		}

		return f
	})

	c.JSON(200, gin.H{
		"accounts": result,
	})
}

func Group(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func Recommend(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func Suggest(c *gin.Context) {
	c.JSON(200, gin.H{})
}
