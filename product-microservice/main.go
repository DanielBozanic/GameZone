package main

import (
	"log"
	"product/config"
	"product/db"
	"product/di"

	"github.com/gin-gonic/gin"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error while loading config file: ", configErr)
	}

	database, databaseError := db.ConnectDatabase(config)


	if databaseError != nil {
		log.Fatal("Error while connecting to database: ", databaseError)
	}

	videoGameAPI := di.InitVideoGameAPI(database)

	r := gin.Default()

	api := r.Group("/api")
	videoGames := api.Group("/videoGames")

	videoGames.GET("", videoGameAPI.GetAll)
	videoGames.GET("/:id", videoGameAPI.GetByID)
	videoGames.GET("/getByName/:name", videoGameAPI.GetByName)
	videoGames.POST("", videoGameAPI.Create)
	videoGames.PUT("", videoGameAPI.Update)
	videoGames.DELETE("/:id", videoGameAPI.Delete)

	err := r.Run(":7000")
	if err != nil {
		panic(err)
	}
}