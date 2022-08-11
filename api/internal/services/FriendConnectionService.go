package services

import (
	"golang_project/api/internal/models"
	"golang_project/api/internal/repositories"
)

type FriendConnectionService interface {
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

func (svc *service) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	var response models.FriendConnectionResponse
	var err error
	response.Success, err = svc.repository.CreateFriendConnection(request.Friends)
	if err != nil {
		panic(err)
	}
	return response
}

func (svc *service) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	var response models.FriendListResponse
	var err error
	response.Friends, err = svc.repository.FindFriendsByEmail(request.Email)
	if err != nil {
		panic(err)
	}
	if response.Friends != nil && len(response.Friends) > 0 {
		response.Success = true
		response.Count = len(response.Friends)
	}
	return response
}

func (svc *service) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	var response models.CommonFriendListResponse
	var err error
	if len(request.Friends) <= 0 {
		return response
	}
	response.Friends, err = svc.repository.FindCommonFriendsByEmails(request.Friends)
	if err != nil {
		panic(err)
	}
	if response.Friends != nil {
		response.Success = true
		response.Count = len(response.Friends)
	}
	return response
}

func (svc *service) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	var response models.SubscribeResponse
	var err error
	response.Success, err = svc.repository.SubscribeFromEmail(request)
	if err != nil {
		panic(err)
	}
	return response
}

func (svc *service) BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	var response models.BlockSubscribeResponse
	var err error
	response.Success, err = svc.repository.BlockSubscribeByEmail(request)
	if err != nil {
		panic(err)
	}
	return response
}

func (svc *service) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	var response models.GetSubscribingEmailListResponse
	var err error
	if request == (models.GetSubscribingEmailListRequest{}) {
		return response
	}
	response.Recipients, err = svc.repository.GetSubscribingEmailListByEmail(request)
	if err != nil {
		panic(err)
	} else {
		response.Success = true
	}
	return response
}
