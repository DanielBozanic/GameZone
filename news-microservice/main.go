package main

import (
	"log"
	"news/config"
	"news/db"
	"news/di"
	"news/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
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

	newsArticleAPI := di.InitNewsArticleAPI(database)
	newsCommentAPI := di.InitNewsCommentAPI(database)
	newsSubscriptionAPI := di.InitNewsSubscriptionAPI(database)

	api := r.Group("/api/news")

	newsArticles := api.Group("/newsArticles")

	newsArticles.GET("/getPublishedArticles", newsArticleAPI.GetPublishedArticles)
	newsArticles.GET("/getNumberOfRecordsPublishedArticles", newsArticleAPI.GetNumberOfRecordsPublishedArticles)
	newsArticles.GET("/:id", newsArticleAPI.GetById)

	employeeProtectedNewsArticles := newsArticles.Group("/employeeProtected")
	employeeProtectedNewsArticles.Use(middleware.AuthorizationRequired([]string { "ROLE_EMPLOYEE" }))
	employeeProtectedNewsArticles.GET("", newsArticleAPI.GetAll)
	employeeProtectedNewsArticles.GET("/getNumberOfRecords", newsArticleAPI.GetNumberOfRecords)
	employeeProtectedNewsArticles.GET("/getUnpublishedArticles", newsArticleAPI.GetUnpublishedArticles)
	employeeProtectedNewsArticles.GET("/getNumberOfRecordsUnpublishedArticles", newsArticleAPI.GetNumberOfRecordsUnpublishedArticles)
	employeeProtectedNewsArticles.POST("/addNewsArticle", newsArticleAPI.AddNewsArticle)
	employeeProtectedNewsArticles.PUT("/editNewsArticle", newsArticleAPI.EditNewsArticle)
	employeeProtectedNewsArticles.DELETE("/deleteNewsArticle/:id", newsArticleAPI.DeleteNewsArticle)
	employeeProtectedNewsArticles.PUT("/publishNewsArticle", newsArticleAPI.PublishNewsArticle)

	newsComments := api.Group("/newsComments")
	newsComments.GET("/getByNewsArticle/:newsArticleId", newsCommentAPI.GetByNewsArticle)

	userProtectedNewsComments := newsComments.Group("/userProtected")
	userProtectedNewsComments.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" }))
	userProtectedNewsComments.POST("/addNewsComment", newsCommentAPI.AddNewsComment)
	userProtectedNewsComments.PUT("/editNewsComment", newsCommentAPI.EditNewsCommment)

	userAndAdminProtectedNewsComments := newsComments.Group("/userAndAdminProtected")
	userAndAdminProtectedNewsComments.Use(middleware.AuthorizationRequired([]string { "ROLE_USER", "ROLE_ADMIN" }))
	userAndAdminProtectedNewsComments.DELETE("/deleteNewsComment/:id", newsCommentAPI.DeleteNewsComment)
	userAndAdminProtectedNewsComments.GET("/getByUserId/:userId", newsCommentAPI.GetByUserId)

	newsSubscriptions := api.Group("/newsSubscriptions")
	
	userProtectedNewsSubscriptions := newsSubscriptions.Group("/userProtected")
	userProtectedNewsSubscriptions.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" }))
	userProtectedNewsSubscriptions.POST("/subscribe", newsSubscriptionAPI.Subscribe)
	userProtectedNewsSubscriptions.DELETE("/unsubscribe", newsSubscriptionAPI.Unsubscribe)
	userProtectedNewsSubscriptions.GET("/isUserSubscribed", newsSubscriptionAPI.IsUserSubscribed)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(7).Days().Do(newsSubscriptionAPI.INewsSubscriptionService.SendEmails)
	scheduler.StartAsync()

	err := r.Run(":7002")
	if err != nil {
		panic(err)
	}
}