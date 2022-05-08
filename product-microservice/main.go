package main

import (
	"log"
	"product/config"
	"product/db"
	"product/di"
	"product/middleware"
	"time"

	"github.com/gin-contrib/cors"
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

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))


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

	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).POST("/addProductToCart/:productId", productAPI.AddProductToCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).GET("/getCurrentCart", productAPI.GetCurrentCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER", "ROLE_ADMIN" })).GET("/getPurchaseHistory/:userId", productAPI.GetPurchaseHistory)
	api.GET("/searchByName", productAPI.SearchByName)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).PUT("/updatePurchase", productAPI.UpdatePurchase)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).DELETE("/removeProductFromCart/:id", productAPI.RemoveProductFromCart)
	api.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" })).PUT("/confirmPurchase", productAPI.ConfirmPurchase)

	videoGames.GET("", videoGameAPI.GetAll)
	videoGames.GET("/:id", videoGameAPI.GetByID)
	videoGames.GET("/searchByName", videoGameAPI.SearchByName)
	videoGames.POST("/filter", videoGameAPI.Filter)
	videoGames.GET("/getPlatforms", videoGameAPI.GetPlatforms)
	videoGames.GET("/getGenres", videoGameAPI.GetGenres)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", videoGameAPI.Create)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", videoGameAPI.Update)
	videoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", videoGameAPI.Delete)

	consoles.GET("", consoleAPI.GetAll)
	consoles.GET("/:id", consoleAPI.GetByID)
	consoles.GET("/searchByName", consoleAPI.SearchByName)
	consoles.POST("/filter", consoleAPI.Filter)
	consoles.GET("/getPlatforms", consoleAPI.GetPlatforms)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", consoleAPI.Create)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", consoleAPI.Update)
	consoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", consoleAPI.Delete)

	graphicsCards.GET("", graphicsCardAPI.GetAll)
	graphicsCards.GET("/:id", graphicsCardAPI.GetByID)
	graphicsCards.GET("/searchByName", graphicsCardAPI.SearchByName)
	graphicsCards.POST("/filter", graphicsCardAPI.Filter)
	graphicsCards.GET("/getManufacturers", graphicsCardAPI.GetManufacturers)
	graphicsCards.GET("/getChipManufacturers", graphicsCardAPI.GetChipManufacturers)
	graphicsCards.GET("/getMemorySizes", graphicsCardAPI.GetMemorySizes)
	graphicsCards.GET("/getMemoryTypes", graphicsCardAPI.GetMemoryTypes)
	graphicsCards.GET("/getModelNames", graphicsCardAPI.GetModelNames)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", graphicsCardAPI.Create)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", graphicsCardAPI.Update)
	graphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", graphicsCardAPI.Delete)

	processors.GET("", processorAPI.GetAll)
	processors.GET("/:id", processorAPI.GetByID)
	processors.GET("/searchByName", processorAPI.SearchByName)
	processors.POST("/filter", processorAPI.Filter)
	processors.GET("/getManufacturers", processorAPI.GetManufacturers)
	processors.GET("/getTypes", processorAPI.GetTypes)
	processors.GET("/getSockets", processorAPI.GetSockets)
	processors.GET("/getNumberOfCores", processorAPI.GetNumberOfCores)
	processors.GET("/getThreads", processorAPI.GetThreads)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", processorAPI.Create)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", processorAPI.Update)
	processors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", processorAPI.Delete)

	motherboards.GET("", motherboardAPI.GetAll)
	motherboards.GET("/:id", motherboardAPI.GetByID)
	motherboards.GET("/searchByName", motherboardAPI.SearchByName)
	motherboards.POST("/filter", motherboardAPI.Filter)
	motherboards.GET("/getManufacturers", motherboardAPI.GetManufacturers)
	motherboards.GET("/getProcessorTypes", motherboardAPI.GetProcessorTypes)
	motherboards.GET("/getSockets", motherboardAPI.GetSockets)
	motherboards.GET("/getFormFactors", motherboardAPI.GetFormFactors)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", motherboardAPI.Create)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", motherboardAPI.Update)
	motherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", motherboardAPI.Delete)

	rams.GET("", ramAPI.GetAll)
	rams.GET("/:id", ramAPI.GetByID)
	rams.GET("/searchByName", ramAPI.SearchByName)
	rams.POST("/filter", ramAPI.Filter)
	rams.GET("/getManufacturers", ramAPI.GetManufacturers)
	rams.GET("/getCapacities", ramAPI.GetCapacities)
	rams.GET("/getMemoryTypes", ramAPI.GetMemoryTypes)
	rams.GET("/getSpeeds", ramAPI.GetSpeeds)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", ramAPI.Create)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", ramAPI.Update)
	rams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", ramAPI.Delete)

	ssds.GET("", solidStateDriveAPI.GetAll)
	ssds.GET("/:id", solidStateDriveAPI.GetByID)
	ssds.GET("/searchByName", solidStateDriveAPI.SearchByName)
	ssds.POST("/filter", solidStateDriveAPI.Filter)
	ssds.GET("/getCapacities", solidStateDriveAPI.GetCapacities)
	ssds.GET("/getForms", solidStateDriveAPI.GetForms)
	ssds.GET("/getManufacturers", solidStateDriveAPI.GetManufacturers)
	ssds.GET("/getMaxSequentialReads", solidStateDriveAPI.GetMaxSequentialReads)
	ssds.GET("/getMaxSequentialWrites", solidStateDriveAPI.GetMaxSequentialWrites)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", solidStateDriveAPI.Create)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", solidStateDriveAPI.Update)
	ssds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", solidStateDriveAPI.Delete)

	hdds.GET("", hardDiskDriveAPI.GetAll)
	hdds.GET("/:id", hardDiskDriveAPI.GetByID)
	hdds.GET("/searchByName", hardDiskDriveAPI.SearchByName)
	hdds.POST("/filter", hardDiskDriveAPI.Filter)
	hdds.GET("/getCapacities", hardDiskDriveAPI.GetCapacities)
	hdds.GET("/getForms", hardDiskDriveAPI.GetForms)
	hdds.GET("/getManufacturers", hardDiskDriveAPI.GetManufacturers)
	hdds.GET("/getDiskSpeeds", hardDiskDriveAPI.GetDiskSpeeds)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", hardDiskDriveAPI.Create)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", hardDiskDriveAPI.Update)
	hdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", hardDiskDriveAPI.Delete)

	monitors.GET("", monitorAPI.GetAll)
	monitors.GET("/:id", monitorAPI.GetByID)
	monitors.GET("/searchByName", monitorAPI.SearchByName)
	monitors.POST("/filter", monitorAPI.Filter)
	monitors.GET("/getManufacturers", monitorAPI.GetManufacturers)
	monitors.GET("/getAspectRatios", monitorAPI.GetAspectRatios)
	monitors.GET("/getResolutions", monitorAPI.GetResolutions)
	monitors.GET("/getRefreshRates", monitorAPI.GetRefreshRates)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", monitorAPI.Create)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", monitorAPI.Update)
	monitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", monitorAPI.Delete)

	psus.GET("", powerSupplyUnitAPI.GetAll)
	psus.GET("/:id", powerSupplyUnitAPI.GetByID)
	psus.GET("/searchByName", powerSupplyUnitAPI.SearchByName)
	psus.POST("/filter", powerSupplyUnitAPI.Filter)
	psus.GET("/getManufacturers", powerSupplyUnitAPI.GetManufacturers)
	psus.GET("/getPowers", powerSupplyUnitAPI.GetPowers)
	psus.GET("/getTypes", powerSupplyUnitAPI.GetTypes)
	psus.GET("/getFormFactors", powerSupplyUnitAPI.GetFormFactors)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", powerSupplyUnitAPI.Create)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", powerSupplyUnitAPI.Update)
	psus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", powerSupplyUnitAPI.Delete)

	keyboards.GET("", keyboardAPI.GetAll)
	keyboards.GET("/:id", keyboardAPI.GetByID)
	keyboards.GET("/searchByName", keyboardAPI.SearchByName)
	keyboards.POST("/filter", keyboardAPI.Filter)
	keyboards.GET("/getManufacturers", keyboardAPI.GetManufacturers)
	keyboards.GET("/getKeyboardConnectors", keyboardAPI.GetKeyboardConnectors)
	keyboards.GET("/getKeyTypes", keyboardAPI.GetKeyTypes)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", keyboardAPI.Create)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", keyboardAPI.Update)
	keyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", keyboardAPI.Delete)

	mouses.GET("", mouseAPI.GetAll)
	mouses.GET("/:id", mouseAPI.GetByID)
	mouses.GET("/searchByName", mouseAPI.SearchByName)
	mouses.POST("/filter", mouseAPI.Filter)
	mouses.GET("/getManufacturers", mouseAPI.GetManufacturers)
	mouses.GET("/getDpis", mouseAPI.GetDPIs)
	mouses.GET("/getConnections", mouseAPI.GetConnections)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", mouseAPI.Create)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", mouseAPI.Update)
	mouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", mouseAPI.Delete)

	headphones.GET("", headphonesAPI.GetAll)
	headphones.GET("/:id", headphonesAPI.GetByID)
	headphones.GET("/searchByName", headphonesAPI.SearchByName)
	headphones.POST("/filter", headphonesAPI.Filter)
	headphones.GET("/getManufacturers", headphonesAPI.GetManufacturers)
	headphones.GET("/getConnectionTypes", headphonesAPI.GetConnectionTypes)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).POST("", headphonesAPI.Create)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).PUT("", headphonesAPI.Update)
	headphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" })).DELETE("/:id", headphonesAPI.Delete)

	err := r.Run(":7000")
	if err != nil {
		panic(err)
	}
}