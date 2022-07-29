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
	  }))


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

	employeeProtectedProducts := api.Group("/employeeProtected")
	employeeProtectedProducts.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE"}))

	employeeProtectedProducts.DELETE("/deleteProduct/:id", productAPI.DeleteProduct)

	// Product purchase API
	productPurchases := api.Group("/productPurchases")

	userProtectedProducts := productPurchases.Group("/userProtected")
	userProtectedProducts.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" }))
	userProtectedProducts.GET("/checkIfProductIsPaidFor", productPurchaseAPI.CheckIfProductIsPaidFor)
	userProtectedProducts.POST("/confirmPurchase", productPurchaseAPI.ConfirmPurchase)
	userProtectedProducts.POST("/sendPurchaseConfirmationMail", productPurchaseAPI.SendPurchaseConfirmationMail)
	userProtectedProducts.GET("/getProductAlertByProductIdAndUserId", productPurchaseAPI.GetProductAlertByProductIdAndUserId)
	userProtectedProducts.POST("/addProductAlert", productPurchaseAPI.AddProductAlert)

	employeeProtectedProductPurchases := productPurchases.Group("/employeeProtected")
	employeeProtectedProductPurchases.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE"}))
	employeeProtectedProductPurchases.GET("/notifyProductAvailability", productPurchaseAPI.NotifyProductAvailability)

	adminAndUserProtectedProductPurchases := productPurchases.Group("/adminAndUserProtected")
	adminAndUserProtectedProductPurchases.Use(middleware.AuthorizationRequired([]string { "ROLE_USER", "ROLE_ADMIN" }))
	adminAndUserProtectedProductPurchases.GET("/getPurchaseHistory", productPurchaseAPI.GetPurchaseHistory)
	adminAndUserProtectedProductPurchases.GET("/getNumberOfRecordsPurchaseHistory", productPurchaseAPI.GetNumberOfRecordsPurchaseHistory)

	adminProtectedProductPurchases := productPurchases.Group("/adminProtected")
	adminProtectedProductPurchases.Use(middleware.AuthorizationRequired([]string { "ROLE_ADMIN"}))
	adminProtectedProductPurchases.PUT("/confirmPayment", productPurchaseAPI.ConfirmPayment)
	adminProtectedProductPurchases.POST("/sendPurchasedDigitalVideoGames", productPurchaseAPI.SendPurchasedDigitalVideoGames)
	
	// Video game API
	employeeProtectedVideoGames := videoGames.Group("/employeeProtected")
	employeeProtectedVideoGames.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedVideoGames.PUT("", videoGameAPI.Update)
	employeeProtectedVideoGames.POST("", videoGameAPI.Create)
	employeeProtectedVideoGames.DELETE("/:id", videoGameAPI.Delete)

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
	employeeProtectedConsoles := consoles.Group("/employeeProtected")
	employeeProtectedConsoles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedConsoles.POST("", consoleAPI.Create)
	employeeProtectedConsoles.PUT("", consoleAPI.Update)
	employeeProtectedConsoles.DELETE("/:id", consoleAPI.Delete)

	consoles.GET("", consoleAPI.GetAll)
	consoles.GET("/getNumberOfRecords", consoleAPI.GetNumberOfRecords)
	consoles.GET("/:id", consoleAPI.GetByID)
	consoles.GET("/searchByName", consoleAPI.SearchByName)
	consoles.GET("/getNumberOfRecordsSearch", consoleAPI.GetNumberOfRecordsSearch)
	consoles.POST("/filter", consoleAPI.Filter)
	consoles.POST("/getNumberOfRecordsFilter", consoleAPI.GetNumberOfRecordsFilter)
	consoles.GET("/getPlatforms", consoleAPI.GetPlatforms)

	// Graphics card API
	employeeProtectedGraphicsCards := graphicsCards.Group("/employeeProtected")
	employeeProtectedGraphicsCards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedGraphicsCards.POST("", graphicsCardAPI.Create)
	employeeProtectedGraphicsCards.PUT("", graphicsCardAPI.Update)
	employeeProtectedGraphicsCards.DELETE("/:id", graphicsCardAPI.Delete)

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
	employeeProtectedProcessors := processors.Group("/employeeProtected")
	employeeProtectedProcessors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedProcessors.POST("", processorAPI.Create)
	employeeProtectedProcessors.PUT("", processorAPI.Update)
	employeeProtectedProcessors.DELETE("/:id", processorAPI.Delete)

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
	employeeProtectedMotherboards := motherboards.Group("/employeeProtected")
	employeeProtectedMotherboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedMotherboards.POST("", motherboardAPI.Create)
	employeeProtectedMotherboards.PUT("", motherboardAPI.Update)
	employeeProtectedMotherboards.DELETE("/:id", motherboardAPI.Delete)

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
	employeeProtectedRams := rams.Group("/employeeProtected")
	employeeProtectedRams.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedRams.POST("", ramAPI.Create)
	employeeProtectedRams.PUT("", ramAPI.Update)
	employeeProtectedRams.DELETE("/:id", ramAPI.Delete)

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
	employeeProtectedSsds := ssds.Group("/employeeProtected")
	employeeProtectedSsds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedSsds.POST("", solidStateDriveAPI.Create)
	employeeProtectedSsds.PUT("", solidStateDriveAPI.Update)
	employeeProtectedSsds.DELETE("/:id", solidStateDriveAPI.Delete)

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
	employeeProtectedHdds := hdds.Group("/employeeProtected")
	employeeProtectedHdds.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedHdds.POST("", hardDiskDriveAPI.Create)
	employeeProtectedHdds.PUT("", hardDiskDriveAPI.Update)
	employeeProtectedHdds.DELETE("/:id", hardDiskDriveAPI.Delete)

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
	employeeProtectedMonitors := monitors.Group("/employeeProtected")
	employeeProtectedMonitors.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedMonitors.POST("", monitorAPI.Create)
	employeeProtectedMonitors.PUT("", monitorAPI.Update)
	employeeProtectedMonitors.DELETE("/:id", monitorAPI.Delete)

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
	employeeProtectedPsus := psus.Group("/employeeProtected")
	employeeProtectedPsus.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedPsus.POST("", powerSupplyUnitAPI.Create)
	employeeProtectedPsus.PUT("", powerSupplyUnitAPI.Update)
	employeeProtectedPsus.DELETE("/:id", powerSupplyUnitAPI.Delete)

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
	employeeProtectedKeyboards := keyboards.Group("/employeeProtected")
	employeeProtectedKeyboards.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedKeyboards.POST("", keyboardAPI.Create)
	employeeProtectedKeyboards.PUT("", keyboardAPI.Update)
	employeeProtectedKeyboards.DELETE("/:id", keyboardAPI.Delete)

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
	employeeProtectedMouses := mouses.Group("/employeeProtected")
	employeeProtectedMouses.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedMouses.POST("", mouseAPI.Create)
	employeeProtectedMouses.PUT("", mouseAPI.Update)
	employeeProtectedMouses.DELETE("/:id", mouseAPI.Delete)

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
	employeeProtectedHeadphones := headphones.Group("/employeeProtected")
	employeeProtectedHeadphones.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedHeadphones.POST("", headphonesAPI.Create)
	employeeProtectedHeadphones.PUT("", headphonesAPI.Update)
	employeeProtectedHeadphones.DELETE("/:id", headphonesAPI.Delete)

	headphones.GET("", headphonesAPI.GetAll)
	headphones.GET("/getNumberOfRecords", mouseAPI.GetNumberOfRecords)
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