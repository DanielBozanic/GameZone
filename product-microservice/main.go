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
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }), middleware.RequestCancelRecover())


	productAPI := di.InitProductAPI(database)
	productPurchaseAPI := di.InitProductPurchaseAPI(database)
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

	// Product API
	api.GET("/:id", productAPI.GetProductById)
	api.GET("/searchByName", productAPI.SearchByName)
	api.GET("/getNumberOfRecordsSearch", productAPI.GetNumberOfRecordsSearch)
	api.GET("/getMainPageProducts", productAPI.GetMainPageProducts)
	api.GET("/getPopularProducts", productAPI.GetPopularProducts)
	api.PUT("/addProductToMainPage/:productId", productAPI.AddProductToMainPage)
	api.PUT("/removeProductFromMainPage/:productId", productAPI.RemoveProductFromMainPage)
	api.DELETE("/deleteProduct/:id", productAPI.DeleteProduct)
	api.GET("/getRecommendedProducts", productAPI.GetRecommendedProducts)

	// Product purchase API
	productPurchases := api.Group("/productPurchases")
	productPurchases.GET("/checkIfProductIsPaidFor", productPurchaseAPI.CheckIfProductIsPaidFor)
	productPurchases.POST("/confirmPurchase", productPurchaseAPI.ConfirmPurchase)
	productPurchases.POST("/sendPurchaseConfirmationMail", productPurchaseAPI.SendPurchaseConfirmationMail)
	productPurchases.GET("/getProductAlertByProductIdAndUserId", productPurchaseAPI.GetProductAlertByProductIdAndUserId)
	productPurchases.POST("/addProductAlert", productPurchaseAPI.AddProductAlert)
	productPurchases.GET("/notifyProductAvailability", productPurchaseAPI.NotifyProductAvailability)
	productPurchases.GET("/getPurchaseHistory", productPurchaseAPI.GetPurchaseHistory)
	productPurchases.GET("/getNumberOfRecordsPurchaseHistory", productPurchaseAPI.GetNumberOfRecordsPurchaseHistory)
	productPurchases.PUT("/confirmPayment", productPurchaseAPI.ConfirmPayment)
	productPurchases.POST("/sendPurchasedDigitalVideoGames", productPurchaseAPI.SendPurchasedDigitalVideoGames)
	
	// Video game API
	videoGames.POST("", videoGameAPI.Create)
	videoGames.PUT("", videoGameAPI.Update)
	videoGames.DELETE("/:id", videoGameAPI.Delete)
	videoGames.GET("", videoGameAPI.GetAll)
	videoGames.GET("/getNumberOfRecords", videoGameAPI.GetNumberOfRecords)
	videoGames.GET("/:id", videoGameAPI.GetByID)
	videoGames.GET("/searchByName", videoGameAPI.SearchByName)
	videoGames.GET("/getNumberOfRecordsSearch", videoGameAPI.GetNumberOfRecordsSearch)
	videoGames.POST("/filter", videoGameAPI.Filter)
	videoGames.POST("/getNumberOfRecordsFilter", videoGameAPI.GetNumberOfRecordsFilter)
	videoGames.GET("/getPlatforms", videoGameAPI.GetPlatforms)
	videoGames.GET("/getGenres", videoGameAPI.GetGenres)

	// Console API
	consoles.POST("", consoleAPI.Create)
	consoles.PUT("", consoleAPI.Update)
	consoles.DELETE("/:id", consoleAPI.Delete)
	consoles.GET("", consoleAPI.GetAll)
	consoles.GET("/getNumberOfRecords", consoleAPI.GetNumberOfRecords)
	consoles.GET("/:id", consoleAPI.GetByID)
	consoles.GET("/searchByName", consoleAPI.SearchByName)
	consoles.GET("/getNumberOfRecordsSearch", consoleAPI.GetNumberOfRecordsSearch)
	consoles.POST("/filter", consoleAPI.Filter)
	consoles.POST("/getNumberOfRecordsFilter", consoleAPI.GetNumberOfRecordsFilter)
	consoles.GET("/getPlatforms", consoleAPI.GetPlatforms)

	// Graphics card API
	graphicsCards.POST("", graphicsCardAPI.Create)
	graphicsCards.PUT("", graphicsCardAPI.Update)
	graphicsCards.DELETE("/:id", graphicsCardAPI.Delete)
	graphicsCards.GET("", graphicsCardAPI.GetAll)
	graphicsCards.GET("/getNumberOfRecords", graphicsCardAPI.GetNumberOfRecords)
	graphicsCards.GET("/:id", graphicsCardAPI.GetByID)
	graphicsCards.GET("/searchByName", graphicsCardAPI.SearchByName)
	graphicsCards.GET("/getNumberOfRecordsSearch", graphicsCardAPI.GetNumberOfRecordsSearch)
	graphicsCards.POST("/filter", graphicsCardAPI.Filter)
	graphicsCards.POST("/getNumberOfRecordsFilter", graphicsCardAPI.GetNumberOfRecordsFilter)
	graphicsCards.GET("/getManufacturers", graphicsCardAPI.GetManufacturers)
	graphicsCards.GET("/getChipManufacturers", graphicsCardAPI.GetChipManufacturers)
	graphicsCards.GET("/getMemorySizes", graphicsCardAPI.GetMemorySizes)
	graphicsCards.GET("/getMemoryTypes", graphicsCardAPI.GetMemoryTypes)
	graphicsCards.GET("/getModelNames", graphicsCardAPI.GetModelNames)

	// Processor API
	processors.POST("", processorAPI.Create)
	processors.PUT("", processorAPI.Update)
	processors.DELETE("/:id", processorAPI.Delete)
	processors.GET("", processorAPI.GetAll)
	processors.GET("/getNumberOfRecords", processorAPI.GetNumberOfRecords)
	processors.GET("/:id", processorAPI.GetByID)
	processors.GET("/searchByName", processorAPI.SearchByName)
	processors.GET("/getNumberOfRecordsSearch", processorAPI.GetNumberOfRecordsSearch)
	processors.POST("/filter", processorAPI.Filter)
	processors.POST("/getNumberOfRecordsFilter", processorAPI.GetNumberOfRecordsFilter)
	processors.GET("/getManufacturers", processorAPI.GetManufacturers)
	processors.GET("/getTypes", processorAPI.GetTypes)
	processors.GET("/getSockets", processorAPI.GetSockets)
	processors.GET("/getNumberOfCores", processorAPI.GetNumberOfCores)
	processors.GET("/getThreads", processorAPI.GetThreads)

	// Motherboard API
	motherboards.POST("", motherboardAPI.Create)
	motherboards.PUT("", motherboardAPI.Update)
	motherboards.DELETE("/:id", motherboardAPI.Delete)
	motherboards.GET("", motherboardAPI.GetAll)
	motherboards.GET("/getNumberOfRecords", motherboardAPI.GetNumberOfRecords)
	motherboards.GET("/:id", motherboardAPI.GetByID)
	motherboards.GET("/searchByName", motherboardAPI.SearchByName)
	motherboards.GET("/getNumberOfRecordsSearch", motherboardAPI.GetNumberOfRecordsSearch)
	motherboards.POST("/filter", motherboardAPI.Filter)
	motherboards.POST("/getNumberOfRecordsFilter", motherboardAPI.GetNumberOfRecordsFilter)
	motherboards.GET("/getManufacturers", motherboardAPI.GetManufacturers)
	motherboards.GET("/getProcessorTypes", motherboardAPI.GetProcessorTypes)
	motherboards.GET("/getSockets", motherboardAPI.GetSockets)
	motherboards.GET("/getFormFactors", motherboardAPI.GetFormFactors)

	// RAM API
	rams.POST("", ramAPI.Create)
	rams.PUT("", ramAPI.Update)
	rams.DELETE("/:id", ramAPI.Delete)
	rams.GET("", ramAPI.GetAll)
	rams.GET("/getNumberOfRecords", ramAPI.GetNumberOfRecords)
	rams.GET("/:id", ramAPI.GetByID)
	rams.GET("/searchByName", ramAPI.SearchByName)
	rams.GET("/getNumberOfRecordsSearch", ramAPI.GetNumberOfRecordsSearch)
	rams.POST("/filter", ramAPI.Filter)
	rams.POST("/getNumberOfRecordsFilter", ramAPI.GetNumberOfRecordsFilter)
	rams.GET("/getManufacturers", ramAPI.GetManufacturers)
	rams.GET("/getCapacities", ramAPI.GetCapacities)
	rams.GET("/getMemoryTypes", ramAPI.GetMemoryTypes)
	rams.GET("/getSpeeds", ramAPI.GetSpeeds)

	// SSD API
	ssds.POST("", solidStateDriveAPI.Create)
	ssds.PUT("", solidStateDriveAPI.Update)
	ssds.DELETE("/:id", solidStateDriveAPI.Delete)
	ssds.GET("", solidStateDriveAPI.GetAll)
	ssds.GET("/getNumberOfRecords", solidStateDriveAPI.GetNumberOfRecords)
	ssds.GET("/:id", solidStateDriveAPI.GetByID)
	ssds.GET("/searchByName", solidStateDriveAPI.SearchByName)
	ssds.GET("/getNumberOfRecordsSearch", solidStateDriveAPI.GetNumberOfRecordsSearch)
	ssds.POST("/filter", solidStateDriveAPI.Filter)
	ssds.POST("/getNumberOfRecordsFilter", solidStateDriveAPI.GetNumberOfRecordsFilter)
	ssds.GET("/getCapacities", solidStateDriveAPI.GetCapacities)
	ssds.GET("/getForms", solidStateDriveAPI.GetForms)
	ssds.GET("/getManufacturers", solidStateDriveAPI.GetManufacturers)
	ssds.GET("/getMaxSequentialReads", solidStateDriveAPI.GetMaxSequentialReads)
	ssds.GET("/getMaxSequentialWrites", solidStateDriveAPI.GetMaxSequentialWrites)

	// HDD API
	hdds.POST("", hardDiskDriveAPI.Create)
	hdds.PUT("", hardDiskDriveAPI.Update)
	hdds.DELETE("/:id", hardDiskDriveAPI.Delete)
	hdds.GET("", hardDiskDriveAPI.GetAll)
	hdds.GET("/getNumberOfRecords", hardDiskDriveAPI.GetNumberOfRecords)
	hdds.GET("/:id", hardDiskDriveAPI.GetByID)
	hdds.GET("/searchByName", hardDiskDriveAPI.SearchByName)
	hdds.GET("/getNumberOfRecordsSearch", hardDiskDriveAPI.GetNumberOfRecordsSearch)
	hdds.POST("/filter", hardDiskDriveAPI.Filter)
	hdds.POST("/getNumberOfRecordsFilter", hardDiskDriveAPI.GetNumberOfRecordsFilter)
	hdds.GET("/getCapacities", hardDiskDriveAPI.GetCapacities)
	hdds.GET("/getForms", hardDiskDriveAPI.GetForms)
	hdds.GET("/getManufacturers", hardDiskDriveAPI.GetManufacturers)
	hdds.GET("/getDiskSpeeds", hardDiskDriveAPI.GetDiskSpeeds)

	// Monitor API
	monitors.POST("", monitorAPI.Create)
	monitors.PUT("", monitorAPI.Update)
	monitors.DELETE("/:id", monitorAPI.Delete)
	monitors.GET("", monitorAPI.GetAll)
	monitors.GET("/getNumberOfRecords", monitorAPI.GetNumberOfRecords)
	monitors.GET("/:id", monitorAPI.GetByID)
	monitors.GET("/searchByName", monitorAPI.SearchByName)
	monitors.GET("/getNumberOfRecordsSearch", monitorAPI.GetNumberOfRecordsSearch)
	monitors.POST("/filter", monitorAPI.Filter)
	monitors.POST("/getNumberOfRecordsFilter", monitorAPI.GetNumberOfRecordsFilter)
	monitors.GET("/getManufacturers", monitorAPI.GetManufacturers)
	monitors.GET("/getAspectRatios", monitorAPI.GetAspectRatios)
	monitors.GET("/getResolutions", monitorAPI.GetResolutions)
	monitors.GET("/getRefreshRates", monitorAPI.GetRefreshRates)

	// PSU API
	psus.POST("", powerSupplyUnitAPI.Create)
	psus.PUT("", powerSupplyUnitAPI.Update)
	psus.DELETE("/:id", powerSupplyUnitAPI.Delete)
	psus.GET("", powerSupplyUnitAPI.GetAll)
	psus.GET("/getNumberOfRecords", powerSupplyUnitAPI.GetNumberOfRecords)
	psus.GET("/:id", powerSupplyUnitAPI.GetByID)
	psus.GET("/searchByName", powerSupplyUnitAPI.SearchByName)
	psus.GET("/getNumberOfRecordsSearch", powerSupplyUnitAPI.GetNumberOfRecordsSearch)
	psus.POST("/filter", powerSupplyUnitAPI.Filter)
	psus.POST("/getNumberOfRecordsFilter", powerSupplyUnitAPI.GetNumberOfRecordsFilter)
	psus.GET("/getManufacturers", powerSupplyUnitAPI.GetManufacturers)
	psus.GET("/getPowers", powerSupplyUnitAPI.GetPowers)
	psus.GET("/getTypes", powerSupplyUnitAPI.GetTypes)
	psus.GET("/getFormFactors", powerSupplyUnitAPI.GetFormFactors)

	// Keyboard API
	keyboards.POST("", keyboardAPI.Create)
	keyboards.PUT("", keyboardAPI.Update)
	keyboards.DELETE("/:id", keyboardAPI.Delete)
	keyboards.GET("", keyboardAPI.GetAll)
	keyboards.GET("/getNumberOfRecords", keyboardAPI.GetNumberOfRecords)
	keyboards.GET("/:id", keyboardAPI.GetByID)
	keyboards.GET("/searchByName", keyboardAPI.SearchByName)
	keyboards.GET("/getNumberOfRecordsSearch", keyboardAPI.GetNumberOfRecordsSearch)
	keyboards.POST("/filter", keyboardAPI.Filter)
	keyboards.POST("/getNumberOfRecordsFilter", keyboardAPI.GetNumberOfRecordsFilter)
	keyboards.GET("/getManufacturers", keyboardAPI.GetManufacturers)
	keyboards.GET("/getKeyboardConnectors", keyboardAPI.GetKeyboardConnectors)
	keyboards.GET("/getKeyTypes", keyboardAPI.GetKeyTypes)

	// Mouse API
	mouses.POST("", mouseAPI.Create)
	mouses.PUT("", mouseAPI.Update)
	mouses.DELETE("/:id", mouseAPI.Delete)
	mouses.GET("", mouseAPI.GetAll)
	mouses.GET("/getNumberOfRecords", mouseAPI.GetNumberOfRecords)
	mouses.GET("/:id", mouseAPI.GetByID)
	mouses.GET("/searchByName", mouseAPI.SearchByName)
	mouses.GET("/getNumberOfRecordsSearch", mouseAPI.GetNumberOfRecordsSearch)
	mouses.POST("/filter", mouseAPI.Filter)
	mouses.POST("/getNumberOfRecordsFilter", mouseAPI.GetNumberOfRecordsFilter)
	mouses.GET("/getManufacturers", mouseAPI.GetManufacturers)
	mouses.GET("/getDpis", mouseAPI.GetDPIs)
	mouses.GET("/getConnections", mouseAPI.GetConnections)

	// Headphones API
	headphones.POST("", headphonesAPI.Create)
	headphones.PUT("", headphonesAPI.Update)
	headphones.DELETE("/:id", headphonesAPI.Delete)
	headphones.GET("", headphonesAPI.GetAll)
	headphones.GET("/getNumberOfRecords", headphonesAPI.GetNumberOfRecords)
	headphones.GET("/:id", headphonesAPI.GetByID)
	headphones.GET("/searchByName", headphonesAPI.SearchByName)
	headphones.GET("/getNumberOfRecordsSearch", headphonesAPI.GetNumberOfRecordsSearch)
	headphones.POST("/filter", headphonesAPI.Filter)
	headphones.POST("/getNumberOfRecordsFilter", headphonesAPI.GetNumberOfRecordsFilter)
	headphones.GET("/getManufacturers", headphonesAPI.GetManufacturers)
	headphones.GET("/getConnectionTypes", headphonesAPI.GetConnectionTypes)
	
	err := r.Run(":7000")
	if err != nil {
		panic(err)
	}
}