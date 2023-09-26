package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":        "AIDEN",
		"aidenMessage": "Would you like to play a game?",
	})
	// c.String(http.StatusOK, "Index")
}
