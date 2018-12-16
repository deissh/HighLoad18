package main

import "github.com/gin-gonic/gin"

func Suggest(c *gin.Context) {

	c.Writer.Header().Set("Transfer-Encoding", "identity")
	c.JSON(200, gin.H{})
}
