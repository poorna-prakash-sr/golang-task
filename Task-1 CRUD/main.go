package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"crud-movies-api/controller/movies"
	"crud-movies-api/utils"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dburl := viper.Get("DBURL").(string)
	release := viper.Get("GIN_MODE").(string)

	fmt.Println(release)

	router := gin.Default()
	dbHandler := utils.Init(dburl)

	movies.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dburl,
		})
	})
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})
	router.Run("127.0.0.1:" + port)
}
