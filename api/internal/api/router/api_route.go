package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang_project/api/internal/config"
	"golang_project/api/internal/controllers"
	"golang_project/api/internal/docs"
	"golang_project/api/internal/repositories"
	"golang_project/api/internal/services"
)

// SetupRouter function used to initilize a router for APIs
// no parameter
// return a pointer of gin.Engine
func SetupRouter() *gin.Engine {
	friendConnectionRepository := repositories.New(config.GetDBInstance())
	friendConnectionService := services.New(friendConnectionRepository)
	friendConnectionController := controllers.New(friendConnectionService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/users/createUser", friendConnectionController.CreateUser)

			v1.POST("/friends/createConnection", friendConnectionController.CreateFriendConnection)

			v1.POST("/friends/showFriendsByEmail", friendConnectionController.GetFriendListByEmail)

			v1.POST("/friends/showCommonFriendList", friendConnectionController.ShowCommonFriendList)

			v1.POST("/friends/subscribeFromEmail", friendConnectionController.SubscribeFromEmail)

			v1.POST("/friends/blockSubscribeByEmail", friendConnectionController.BlockSubscribeByEmail)

			v1.POST("/friends/showSubscribingEmailListByEmail", friendConnectionController.GetSubscribingEmailListByEmail)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
