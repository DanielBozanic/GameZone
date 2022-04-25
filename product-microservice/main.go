package main

import (
	"log"
	"product/config"
	"product/db"
	"product/di"
	"product/middleware"

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

	productAPI := di.InitProductAPI(database)
	videoGameAPI := di.InitVideoGameAPI(database)
	consoleAPI := di.InitConsoleAPI(database)
	graphicsCardAPI := di.InitGraphicsCardAPI(database)
	processorAPI := di.InitProcessorAPI(database)
	motherboardAPI := di.InitMotherboardAPI(database)
	ramAPI := di.InitRamAPI(database)
	solidStateDriveAPI := di.InitSolidStateDriveAPI(database)
	hardDiskDriveAPI := di.InitHardDiskDriveAPI(database)
	monitorAPI := di.InitMonitorAPI(database)
	powerSupplyUnitAPI := di.InitPowerSupplyUnitAPI(database)
	keyboardAPI := di.InitKeyboardAPI(database)
	mouseAPI := di.InitMouseAPI(database)
	headphonesAPI := di.InitHeadphonesAPI(database)

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
	monitors := api.Group("/monitors")
	psus := api.Group("/psus")
	keyboards := api.Group("/keyboards")
	mouses := api.Group("/mouses")
	headphones := api.Group("/headphones")

	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).POST("/addToCart", productAPI.AddProductToCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).GET("/getCurrentCart/:userId", productAPI.GetCurrentCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).GET("/getPurchaseHistory/:userId", productAPI.GetPurchaseHistory)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).PUT("/updatePurchase", productAPI.UpdatePurchase)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).DELETE("/removeProductFromCart/:id", productAPI.RemoveProductFromCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).PUT("/confirmPurchase/:userId", productAPI.ConfirmPurchase)

	videoGames.GET("", videoGameAPI.GetAll)
	videoGames.GET("/:id", videoGameAPI.GetByID)
	videoGames.GET("/getByName/:name", videoGameAPI.GetByName)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", videoGameAPI.Create)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", videoGameAPI.Update)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", videoGameAPI.Delete)

	consoles.GET("", consoleAPI.GetAll)
	consoles.GET("/:id", consoleAPI.GetByID)
	consoles.GET("/getByName/:name", consoleAPI.GetByName)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", consoleAPI.Create)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", consoleAPI.Update)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", consoleAPI.Delete)

	graphicsCards.GET("", graphicsCardAPI.GetAll)
	graphicsCards.GET("/:id", graphicsCardAPI.GetByID)
	graphicsCards.GET("/getByName/:name", graphicsCardAPI.GetByName)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", graphicsCardAPI.Create)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", graphicsCardAPI.Update)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", graphicsCardAPI.Delete)

	processors.GET("", processorAPI.GetAll)
	processors.GET("/:id", processorAPI.GetByID)
	processors.GET("/getByName/:name", processorAPI.GetByName)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", processorAPI.Create)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", processorAPI.Update)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", processorAPI.Delete)

	motherboards.GET("", motherboardAPI.GetAll)
	motherboards.GET("/:id", motherboardAPI.GetByID)
	motherboards.GET("/getByName/:name", motherboardAPI.GetByName)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", motherboardAPI.Create)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", motherboardAPI.Update)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", motherboardAPI.Delete)

	rams.GET("", ramAPI.GetAll)
	rams.GET("/:id", ramAPI.GetByID)
	rams.GET("/getByName/:name", ramAPI.GetByName)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", ramAPI.Create)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", ramAPI.Update)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", ramAPI.Delete)

	ssds.GET("", solidStateDriveAPI.GetAll)
	ssds.GET("/:id", solidStateDriveAPI.GetByID)
	ssds.GET("/getByName/:name", solidStateDriveAPI.GetByName)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", solidStateDriveAPI.Create)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", solidStateDriveAPI.Update)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", solidStateDriveAPI.Delete)

	hdds.GET("", hardDiskDriveAPI.GetAll)
	hdds.GET("/:id", hardDiskDriveAPI.GetByID)
	hdds.GET("/getByName/:name", hardDiskDriveAPI.GetByName)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", hardDiskDriveAPI.Create)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", hardDiskDriveAPI.Update)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", hardDiskDriveAPI.Delete)

	monitors.GET("", monitorAPI.GetAll)
	monitors.GET("/:id", monitorAPI.GetByID)
	monitors.GET("/getByName/:name", monitorAPI.GetByName)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", monitorAPI.Create)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", monitorAPI.Update)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", monitorAPI.Delete)

	psus.GET("", powerSupplyUnitAPI.GetAll)
	psus.GET("/:id", powerSupplyUnitAPI.GetByID)
	psus.GET("/getByName/:name", powerSupplyUnitAPI.GetByName)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", powerSupplyUnitAPI.Create)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", powerSupplyUnitAPI.Update)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", powerSupplyUnitAPI.Delete)

	keyboards.GET("", keyboardAPI.GetAll)
	keyboards.GET("/:id", keyboardAPI.GetByID)
	keyboards.GET("/getByName/:name", keyboardAPI.GetByName)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", keyboardAPI.Create)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", keyboardAPI.Update)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", keyboardAPI.Delete)

	mouses.GET("", mouseAPI.GetAll)
	mouses.GET("/:id", mouseAPI.GetByID)
	mouses.GET("/getByName/:name", mouseAPI.GetByName)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", mouseAPI.Create)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", mouseAPI.Update)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", mouseAPI.Delete)

	headphones.GET("", headphonesAPI.GetAll)
	headphones.GET("/:id", headphonesAPI.GetByID)
	headphones.GET("/getByName/:name", headphonesAPI.GetByName)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", headphonesAPI.Create)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", headphonesAPI.Update)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", headphonesAPI.Delete)

	err := r.Run(":7000")
	if err != nil {
		panic(err)
	}
}