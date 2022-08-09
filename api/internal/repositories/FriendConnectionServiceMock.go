package repositories

import (
	"golang_project/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type FriendConnectionServiceMock struct {
	mock.Mock
}

func (s *FriendConnectionServiceMock) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	if len(request.Friends) > 0 {
		return models.FriendConnectionResponse{Success: true}
	}
	return models.FriendConnectionResponse{Success: false}
}
func (s *FriendConnectionServiceMock) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	if request.Email != "" {
		return models.FriendListResponse{Success: false}
	}
	return models.FriendListResponse{Success: true, Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}, Count: 2}
}
func (s *FriendConnectionServiceMock) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	return models.CommonFriendListResponse{}
}
func (s *FriendConnectionServiceMock) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	return models.SubscribeResponse{}
}
func (s *FriendConnectionServiceMock) BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	return models.BlockSubscribeResponse{}
}
func (s *FriendConnectionServiceMock) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	return models.GetSubscribingEmailListResponse{}
}
