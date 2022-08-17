package router

import (
	"golang_project/api/cmd/golang_project/database"
	"golang_project/api/internal/controllers"
	"golang_project/api/internal/docs"
	"golang_project/api/internal/repositories"
	"golang_project/api/internal/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	//init variables
	friendConnectionRepository := repositories.New(database.GetInstance())
	friendConnectionService := services.New(friendConnectionRepository)
	friendConnectionController := controllers.New(friendConnectionService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//create user
			v1.POST("/users/createUser", friendConnectionController.CreateUser)

			//1. Done
			v1.POST("/friends/createConnection", friendConnectionController.CreateFriendConnection)

			//2. Done
			v1.POST("/friends/showFriendsByEmail", friendConnectionController.GetFriendListByEmail)

			//3. Done
			v1.POST("/friends/showCommonFriendList", friendConnectionController.ShowCommonFriendList)

			//4. Done
			v1.POST("/friends/subscribeFromEmail", friendConnectionController.SubscribeFromEmail)

			//5. Done
			v1.POST("/friends/blockSubscribeByEmail", friendConnectionController.BlockSubscribeByEmail)

			//6. Done
			v1.POST("/friends/showSubscribingEmailListByEmail", friendConnectionController.GetSubscribingEmailListByEmail)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
