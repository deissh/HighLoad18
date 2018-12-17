package main

import (
	"github.com/deissh/HighLoad18/core/account"
	filtering "github.com/deissh/HighLoad18/core/filters"
	"github.com/deissh/HighLoad18/core/storage/inmemory"
	rest "github.com/deissh/HighLoad18/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var accounts account.Accounts

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	LoadFromZip(os.Getenv("ZIP_FILE"))

	log.Println("create http server")

	r := gin.New()
	g := r.Group("/accounts")
	{
		g.GET("/filter")
	}

	repository := inmemory.New(accounts)
	filteringService := filtering.New(repository)
	hdl := rest.New(filteringService)

	log.Println("starting server ...")
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), hdl))
}
