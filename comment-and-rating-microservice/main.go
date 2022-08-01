package main

import (
	"comment-and-rating/config"
	"comment-and-rating/db"
	"comment-and-rating/di"
	"comment-and-rating/middleware"
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

	productCommentAPI := di.InitProductCommentAPI(database)

	api := r.Group("/api/comments")

	productComments := api.Group("/productComments")
	productComments.GET("", productCommentAPI.GetAll)
	productComments.GET("/:id", productCommentAPI.GetById)
	productComments.GET("/getByProductId/:productId", productCommentAPI.GetByProductId)
	productComments.GET("/getByProductAndUser", productCommentAPI.GetByProductAndUser)

	userProtectedProductComments := productComments.Group("/userProtected")
	userProtectedProductComments.Use(middleware.AuthorizationRequired([]string { "ROLE_USER" }))
	userProtectedProductComments.POST("/addComment", productCommentAPI.AddComment)
	userProtectedProductComments.PUT("/editComment", productCommentAPI.EditComment)
	
	userAndAdminProtectedProductComments := productComments.Group("/userAndAdminProtected")
	userAndAdminProtectedProductComments.Use(middleware.AuthorizationRequired([]string { "ROLE_USER", "ROLE_ADMIN" }))
	userAndAdminProtectedProductComments.DELETE("/deleteComment/:id", productCommentAPI.DeleteComment)
	userAndAdminProtectedProductComments.GET("/getByUserId/:userId", productCommentAPI.GetByUserId)
	
	err := r.Run(":7001")
	if err != nil {
		panic(err)
	}
}