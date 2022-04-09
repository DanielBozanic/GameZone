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
	consoleAPI := di.InitConsoleAPI(database)
	graphicsCardAPI := di.InitGraphicsCardAPI(database)
	processorAPI := di.InitProcessorAPI(database)
	motherboardAPI := di.InitMotherboardAPI(database)
	ramAPI := di.InitRamAPI(database)
	solidStateDriveAPI := di.InitSoldiStateDriveAPI(database)
	hardDiskDriveAPI := di.InitHardDiskDriveAPI(database)

	r := gin.Default()

	api := r.Group("/api/products")
	videoGames := api.Group("/videoGames")
	consoles := api.Group("/consoles")
	graphicsCards := api.Group("/graphicsCards")
	processors := api.Group("/processors")
	motherboards := api.Group("/motherboards")
	rams := api.Group("/rams")
	ssds := api.Group("/ssds")
	hdds := api.Group("/hdds")

	videoGames.GET("", videoGameAPI.GetAll)
	videoGames.GET("/:id", videoGameAPI.GetByID)
	videoGames.GET("/getByName/:name", videoGameAPI.GetByName)
	videoGames.POST("", videoGameAPI.Create)
	videoGames.PUT("", videoGameAPI.Update)
	videoGames.DELETE("/:id", videoGameAPI.Delete)

	consoles.GET("", consoleAPI.GetAll)
	consoles.GET("/:id", consoleAPI.GetByID)
	consoles.GET("/getByName/:name", consoleAPI.GetByName)
	consoles.POST("", consoleAPI.Create)
	consoles.PUT("", consoleAPI.Update)
	consoles.DELETE("/:id", consoleAPI.Delete)

	graphicsCards.GET("", graphicsCardAPI.GetAll)
	graphicsCards.GET("/:id", graphicsCardAPI.GetByID)
	graphicsCards.GET("/getByName/:name", graphicsCardAPI.GetByName)
	graphicsCards.POST("", graphicsCardAPI.Create)
	graphicsCards.PUT("", graphicsCardAPI.Update)
	graphicsCards.DELETE("/:id", graphicsCardAPI.Delete)

	processors.GET("", processorAPI.GetAll)
	processors.GET("/:id", processorAPI.GetByID)
	processors.GET("/getByName/:name", processorAPI.GetByName)
	processors.POST("", processorAPI.Create)
	processors.PUT("", processorAPI.Update)
	processors.DELETE("/:id", processorAPI.Delete)

	motherboards.GET("", motherboardAPI.GetAll)
	motherboards.GET("/:id", motherboardAPI.GetByID)
	motherboards.GET("/getByName/:name", motherboardAPI.GetByName)
	motherboards.POST("", motherboardAPI.Create)
	motherboards.PUT("", motherboardAPI.Update)
	motherboards.DELETE("/:id", motherboardAPI.Delete)

	rams.GET("", ramAPI.GetAll)
	rams.GET("/:id", ramAPI.GetByID)
	rams.GET("/getByName/:name", ramAPI.GetByName)
	rams.POST("", ramAPI.Create)
	rams.PUT("", ramAPI.Update)
	rams.DELETE("/:id", ramAPI.Delete)

	ssds.GET("", solidStateDriveAPI.GetAll)
	ssds.GET("/:id", solidStateDriveAPI.GetByID)
	ssds.GET("/getByName/:name", solidStateDriveAPI.GetByName)
	ssds.POST("", solidStateDriveAPI.Create)
	ssds.PUT("", solidStateDriveAPI.Update)
	ssds.DELETE("/:id", solidStateDriveAPI.Delete)

	hdds.GET("", hardDiskDriveAPI.GetAll)
	hdds.GET("/:id", hardDiskDriveAPI.GetByID)
	hdds.GET("/getByName/:name", hardDiskDriveAPI.GetByName)
	hdds.POST("", hardDiskDriveAPI.Create)
	hdds.PUT("", hardDiskDriveAPI.Update)
	hdds.DELETE("/:id", hardDiskDriveAPI.Delete)

	err := r.Run(":7000")
	if err != nil {
		panic(err)
	}
}