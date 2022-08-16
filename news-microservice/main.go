package main

import (
	"log"
	"news/config"
	"news/db"
	"news/di"
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
	newsArticles.GET("", newsArticleAPI.GetAll)
	newsArticles.GET("/getNumberOfRecords", newsArticleAPI.GetNumberOfRecords)
	newsArticles.POST("/addNewsArticle", newsArticleAPI.AddNewsArticle)
	newsArticles.PUT("/editNewsArticle", newsArticleAPI.EditNewsArticle)
	newsArticles.DELETE("/deleteNewsArticle/:id", newsArticleAPI.DeleteNewsArticle)
	newsArticles.PUT("/publishNewsArticle", newsArticleAPI.PublishNewsArticle)

	newsComments := api.Group("/newsComments")
	newsComments.GET("/getByNewsArticle/:newsArticleId", newsCommentAPI.GetByNewsArticle)
	newsComments.POST("/addNewsComment", newsCommentAPI.AddNewsComment)
	newsComments.PUT("/editNewsComment", newsCommentAPI.EditNewsCommment)
	newsComments.DELETE("/deleteNewsComment/:id", newsCommentAPI.DeleteNewsComment)
	newsComments.DELETE("/deleteNewsCommentsByNewsArticleId/:newsArticleId", newsCommentAPI.DeleteNewsCommentsByNewsArticleId)
	newsComments.GET("/getByUserId/:userId", newsCommentAPI.GetByUserId)

	newsSubscriptions := api.Group("/newsSubscriptions")
	newsSubscriptions.POST("/subscribe", newsSubscriptionAPI.Subscribe)
	newsSubscriptions.DELETE("/unsubscribe", newsSubscriptionAPI.Unsubscribe)
	newsSubscriptions.GET("/isUserSubscribed", newsSubscriptionAPI.IsUserSubscribed)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(7).Days().Do(newsSubscriptionAPI.INewsSubscriptionService.SendEmails)
	scheduler.StartAsync()

	err := r.Run(":7002")
	if err != nil {
		panic(err)
	}
}