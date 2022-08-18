package services

import (
	"golang_project/api/internal/models"
	"golang_project/api/internal/repositories"
)

type FriendConnectionService interface {
	CreateUser(models.CreatingUserRequest) models.CreatingUserResponse
	CreateConnection(models.FriendConnectionRequest) models.FriendConnectionResponse
	GetFriendConnection(request models.FriendListRequest) models.FriendListResponse
	ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse
	SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse
	BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse
	GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse
}

type service struct {
	repository repositories.FriendConnectionRepository
}

func New(repo repositories.FriendConnectionRepository) FriendConnectionService {
	return &service{
		repository: repo,
	}
}

func (svc *service) CreateUser(request models.CreatingUserRequest) models.CreatingUserResponse {
	response := models.CreatingUserResponse{Success: false}
	var err error
	user, err := svc.repository.CreateUser(request)
	if err != nil {
		panic(err)
	}
	if user != (models.User{}) {
		response.Success = true
	}
	return response
}

func (svc *service) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	var err error
	model, err := svc.repository.CreateFriendConnection(request)
	if err != nil {
		panic(err)
	}
	if (model != models.Relationship{} && model.Is_friend == true) {
		return models.FriendConnectionResponse{Success: true}
	}
	return models.FriendConnectionResponse{Success: false}
}

func (svc *service) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	var err error
	relationships, err := svc.repository.FindFriendsByEmail(request)
	if err != nil {
		panic(err)
	}
	if len(relationships) > 0 {
		var friends []string
		for _, relationship := range relationships {
			friends = append(friends, relationship.Target)
		}
		return models.FriendListResponse{Success: true, Friends: friends, Count: len(friends)}
	}
	return models.FriendListResponse{Success: false}
}

func (svc *service) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	var err error
	if len(request.Friends) == 0 {
		return models.CommonFriendListResponse{Success: false}
	}
	relationships, err := svc.repository.FindCommonFriendsByEmails(request)
	if err != nil {
		panic(err)
	}
	var friends []string
	for _, relationship := range relationships {
		friends = append(friends, relationship.Target)
	}
	return models.CommonFriendListResponse{Success: true, Friends: friends, Count: len(friends)}

}

func (svc *service) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	relationships, err := svc.repository.SubscribeFromEmail(request)
	if err != nil {
		panic(err)
	}
	if (relationships != models.Relationship{}) {
		return models.SubscribeResponse{Success: true}
	}
	return models.SubscribeResponse{Success: false}
}

func (svc *service) BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	relationship, err := svc.repository.BlockSubscribeByEmail(request)
	if err != nil {
		panic(err)
	}
	if relationship.Friend_blocked == true {
		return models.BlockSubscribeResponse{Success: true}
	}
	return models.BlockSubscribeResponse{Success: false}
}

func (svc *service) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	response := models.GetSubscribingEmailListResponse{Success: false}
	var err error
	if request == (models.GetSubscribingEmailListRequest{}) {
		return response
	}
	relationship, err := svc.repository.GetSubscribingEmailListByEmail(request)
	if err != nil {
		panic(err)
	} else {
		response.Success = true
		var recipients []string
		for _, re := range relationship {
			recipients = append(recipients, re.Target)
		}
		response.Recipients = recipients
	}
	return response
}
