package main

import (
	"DistributedAI/checkers"
	"DistributedAI/tictactoe"
	"github.com/gin-gonic/gin"
)

func init() {
	// generateStaticFiles()
}

func main() {
	r := gin.Default()
	_ = r.SetTrustedProxies([]string{})

	r.LoadHTMLGlob("*/*.html")

	r.Static("/static", "./web/")

	r.GET("/", Index)

	ttt := r.Group("/tictactoe")
	ttt.GET("/", tictactoe.Index)
	ttt.GET("/settings", tictactoe.Settings)
	ttt.PUT("/move", tictactoe.ClickMove)

	c := r.Group("/checkers")
	c.GET("/", checkers.Index)
	c.GET("/settings", checkers.Settings)

	_ = r.Run(":8080")
}
