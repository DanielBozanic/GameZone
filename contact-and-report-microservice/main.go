package main

import (
	"contact-and-report/config"
	"contact-and-report/db"
	"contact-and-report/di"
	"contact-and-report/middleware"
	"log"
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

	reportAPI := di.InitReportAPI(database)
	banAPI := di.InitBanAPI(database)
	contactAPI := di.InitContactAPI(database)

	api := r.Group("/api/contactAndReport")

	reports := api.Group("/reports")

	userAndEmployeeProtectedReports := reports.Group("/userAndEmployeeProtected")
	userAndEmployeeProtectedReports.Use(middleware.AuthorizationRequired([]string { "ROLE_USER", "ROLE_EMPLOYEE" }))
	userAndEmployeeProtectedReports.POST("/addReport", reportAPI.AddReport)

	adminProtectedReports := reports.Group("/adminProtected")
	adminProtectedReports.Use(middleware.AuthorizationRequired([]string { "ROLE_ADMIN" }))
	adminProtectedReports.GET("/getReportsByUserId/:userId", reportAPI.GetReportsByUserId)

	bans := api.Group("/bans")

	bans.GET("/isUserBanned/:userId", banAPI.IsUserBanned)

	adminProtectedBans := bans.Group("/adminProtected")
	adminProtectedBans.Use(middleware.AuthorizationRequired([]string { "ROLE_ADMIN" }))
	adminProtectedBans.GET("/getUserBanHistory/:userId", banAPI.GetUserBanHistory)
	adminProtectedBans.POST("/addBan", banAPI.AddBan)
	adminProtectedBans.POST("/sendEmailToBannedUser", banAPI.SendEmailToBannedUser)

	contacts := api.Group("/contacts")

	employeeProtectedContacts := contacts.Group("/employeeProtected")
	employeeProtectedContacts.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedContacts.GET("/getUnansweredContactMessages", contactAPI.GetUnansweredContactMessages)

	adminProtectedContacts := contacts.Group("/adminProtected")
	adminProtectedContacts.Use(middleware.AuthorizationRequired([]string { "ROLE_ADMIN" }))
	adminProtectedContacts.GET("/getUnansweredContactMessagesByUserId/:userId", contactAPI.GetUnansweredContactMessagesByUserId)

	adminAndEmployeeProtectedContacts := contacts.Group("/adminAndEmployeeProtected")
	adminAndEmployeeProtectedContacts.Use(middleware.AuthorizationRequired([]string { "ROLE_ADMIN", "ROLE_EMPLOYEE" }))
	adminAndEmployeeProtectedContacts.PUT("/answerContactMessage", contactAPI.AnswerContactMessage)

	userProtectedContacts := contacts.Group("/userProtected")
	userProtectedContacts.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" }))
	userProtectedContacts.GET("/getContactMessagesByUserId/:userId", contactAPI.GetContactMessagesByUserId)
	userProtectedContacts.POST("/sendContactMessage", contactAPI.SendContactMessage)

	err := r.Run(":7003")
	if err != nil {
		panic(err)
	}
}