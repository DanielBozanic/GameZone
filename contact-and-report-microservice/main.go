package main

import (
	"contact-and-report/config"
	"contact-and-report/db"
	"contact-and-report/di"
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
	reports.POST("/addReport", reportAPI.AddReport)
	reports.GET("/getReportsByUserId/:userId", reportAPI.GetReportsByUserId)

	bans := api.Group("/bans")

	bans.GET("/isUserBanned/:userId", banAPI.IsUserBanned)
	bans.GET("/getUserBanHistory/:userId", banAPI.GetUserBanHistory)
	bans.POST("/addBan", banAPI.AddBan)
	bans.POST("/sendEmailToBannedUser", banAPI.SendEmailToBannedUser)

	contacts := api.Group("/contacts")
	contacts.GET("/getUnansweredContactMessages", contactAPI.GetUnansweredContactMessages)
	contacts.GET("/getUnansweredContactMessagesByUserId/:userId", contactAPI.GetUnansweredContactMessagesByUserId)
	contacts.PUT("/answerContactMessage", contactAPI.AnswerContactMessage)
	contacts.GET("/getContactMessagesByUserId/:userId", contactAPI.GetContactMessagesByUserId)
	contacts.POST("/sendContactMessage", contactAPI.SendContactMessage)

	err := r.Run(":7003")
	if err != nil {
		panic(err)
	}
}