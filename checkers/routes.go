package checkers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "checkers/index.html", gin.H{
		"title": "Checkers",
	})
}

func Settings(c *gin.Context) {
	c.String(http.StatusOK, "Checkers Settings")
}
