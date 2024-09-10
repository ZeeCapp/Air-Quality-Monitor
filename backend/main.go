package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type measurment struct {
	Id               int64     `json:"id"`
	DeviceUUID       string    `json:"deviceUUID"`
	MeasuredDateTime time.Time `json:"measuredDateTime"`
	MeasuredValue    string    `json:"measuredValue"`
}

func main() {
	dbConnString := os.Getenv("DB_CONN_STRING")

	conn, err := pgx.Connect(context.Background(), dbConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	api := gin.Default()

	api.GET("/ping", func(ctx *gin.Context) {
		var greeting string

		err = conn.QueryRow(context.Background(), "select 'pong'").Scan(&greeting)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": greeting,
		})
	})

	api.POST("/measurment", func(ctx *gin.Context) {
		var measurment measurment

		if err := ctx.BindJSON(&measurment); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		}

		measurment.MeasuredDateTime = time.Now()

		var id int64

		err := conn.QueryRow(
			context.Background(),
			"insert into measurments(device_uuid , measured_datetime, measured_value) values ($1, $2, $3) returning id;",
			measurment.DeviceUUID, measurment.MeasuredDateTime, measurment.MeasuredValue,
		).Scan(&id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	api.Run()
}
