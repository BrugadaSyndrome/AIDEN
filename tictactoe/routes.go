package tictactoe

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ttt *TicTacToe
)

func init() {
	ttt = &TicTacToe{}
}

func Index(c *gin.Context) {
	ttt.NewGame()
	c.HTML(http.StatusOK, "tictactoe/index.html", gin.H{
		"title":            "TicTacToe",
		"aidenMessage":     "It is your turn",
		"boardStrokeColor": "#000000",
		"boardStrokeWidth": 2,
		"xStrokeColor":     "#FF0000",
		"xStrokeWidth":     2,
		"oStrokeColor":     "#0000FF",
		"oStrokeWidth":     2,
	})
}

func Settings(c *gin.Context) {
	c.String(http.StatusOK, "TicTacToe Settings")
}

func ClickMove(c *gin.Context) {
	var move Move
	c.BindJSON(&move)
	move = ttt.Move(move)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, move)
}
