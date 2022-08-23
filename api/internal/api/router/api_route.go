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

// SetupRouter function used to initialize a router for APIs
// no parameter
// return a pointer of gin.Engine
func SetupRouter() *gin.Engine {
	friendConnectionRepo := repositories.New(config.GetDBInstance())
	friendConnectionSrv := services.New(friendConnectionRepo)
	friendConnectionCtrl := controllers.New(friendConnectionSrv)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/users/createUser", friendConnectionCtrl.CreateUser)

			v1.POST("/friends/createConnection", friendConnectionCtrl.CreateFriendConnection)

			v1.POST("/friends/showFriendsByEmail", friendConnectionCtrl.GetFriendListByEmail)

			v1.POST("/friends/showCommonFriendList", friendConnectionCtrl.ShowCommonFriendList)

			v1.POST("/friends/subscribeFromEmail", friendConnectionCtrl.SubscribeFromEmail)

			v1.POST("/friends/blockSubscribeByEmail", friendConnectionCtrl.BlockSubscribeByEmail)

			v1.POST("/friends/showSubscribingEmailListByEmail", friendConnectionCtrl.GetSubscribingEmailListByEmail)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
