package test

import (
	"golang_project/controllers"
	"golang_project/docs"
	"golang_project/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	if len(request.Friends) > 0 {
		return models.FriendConnectionResponse{Success: true}
	}
	return models.FriendConnectionResponse{Success: false}
}
func (s *serviceMock) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	if request.Email != "" {
		return models.FriendListResponse{Success: false}
	}
	return models.FriendListResponse{Success: true, Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}, Count: 2}
}
func (s *serviceMock) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	if len(request.Friends)>0{
		return models.CommonFriendListResponse{Success: true, }
	}
	return models.CommonFriendListResponse{}
}
func (s *serviceMock) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	if request.Requestor != "" && request.Target != "" {
		return models.SubscribeResponse{Success: true}
	}
	return models.SubscribeResponse{}
}
func (s *serviceMock) BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	if request.Requestor != "" && request.Target != "" {
		return models.BlockSubscribeResponse{Success: true}
	}
	return models.BlockSubscribeResponse{}
}
func (s *serviceMock) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	if request.Sender != "" && request.Text != "" {
		return models.GetSubscribingEmailListResponse{Success: true, Recipients: []string{"abc@gmail.com"}}
	}
	return models.GetSubscribingEmailListResponse{}
}

func SetupRouterForTesting() *gin.Engine {
	serv := &serviceMock{}
	controller := controllers.New(serv)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//1. Done
			v1.POST("/friends/createConnection", controller.CreateFriendConnection)

			//2. Done
			v1.POST("/friends/showFriendsByEmail", controller.GetFriendListByEmail)

			//3. Done
			v1.POST("/friends/showCommonFriendList", controller.ShowCommonFriendList)

			//4. Done
			v1.POST("/friends/subscribeFromEmail", controller.SubscribeFromEmail)

			//5. Done
			v1.POST("/friends/blockSubscribeByEmail", controller.BlockSubscribeByEmail)

			//6. Done
			v1.POST("/friends/showSubscribingEmailListByEmail", controller.GetSubscribingEmailListByEmail)
		}
	}
	return router
}
