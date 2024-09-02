package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func main() {
	db := pg.Connect(&pg.Options{})

	db.Begin()

	api := gin.Default()

	api.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.Run()
}
