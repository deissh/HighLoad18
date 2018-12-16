package main

import "github.com/gin-gonic/gin"

func Group(c *gin.Context) {
	var keys GroupKey
	var Param Account

	err := c.Bind(&keys)
	if err != nil {
		c.Status(400)
		return
	}

	err = c.Bind(&Param)
	if err != nil {
		c.Status(400)
		return
	}

	c.Writer.Header().Set("Transfer-Encoding", "identity")
	c.JSON(200, gin.H{})
}
