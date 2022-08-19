package services

import (
	"errors"
	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"
	"golang_project/api/internal/repositories"
)

type FriendConnectionService interface {
	CreateUser(models.CreatingUserRequest) (models.CreatingUserResponse, error)
	CreateConnection(models.FriendConnectionRequest) (models.FriendConnectionResponse, error)
	GetFriendConnection(request models.FriendListRequest) (models.FriendListResponse, error)
	ShowCommonFriendList(request models.CommonFriendListRequest) (models.CommonFriendListResponse, error)
	SubscribeFromEmail(request models.SubscribeRequest) (models.SubscribeResponse, error)
	BlockSubscribeByEmail(request models.BlockSubscribeRequest) (models.BlockSubscribeResponse, error)
	GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) (models.GetSubscribingEmailListResponse, error)
}

type service struct {
	repository repositories.FriendConnectionRepository
}

func New(repo repositories.FriendConnectionRepository) FriendConnectionService {
	return &service{
		repository: repo,
	}
}

func (svc *service) CreateUser(request models.CreatingUserRequest) (models.CreatingUserResponse, error) {
	response := models.CreatingUserResponse{Success: false}
	_, err := svc.repository.CreateUser(request)
	if err != nil {
		return models.CreatingUserResponse{}, err
	} else {
		response.Success = true
	}
	return response, nil
}

func (svc *service) CreateConnection(request models.FriendConnectionRequest) (models.FriendConnectionResponse, error) {
	_, err := svc.repository.CreateFriendConnection(request)
	if err != nil {
		return models.FriendConnectionResponse{}, err
	}
	return models.FriendConnectionResponse{Success: true}, nil
}

func (svc *service) GetFriendConnection(request models.FriendListRequest) (models.FriendListResponse, error) {
	relationships, err := svc.repository.FindFriendsByEmail(request)
	if err != nil {
		return models.FriendListResponse{}, err
	}
	if len(relationships) > 0 {
		var friends []string
		for _, relationship := range relationships {
			friends = append(friends, relationship.Target)
		}
		return models.FriendListResponse{Success: true, Friends: friends, Count: len(friends)}, nil
	}
	return models.FriendListResponse{Success: false}, nil
}

func (svc *service) ShowCommonFriendList(request models.CommonFriendListRequest) (models.CommonFriendListResponse, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return models.CommonFriendListResponse{}, err
	}
	relationships, err := svc.repository.FindCommonFriendsByEmails(request)
	if err != nil {
		return models.CommonFriendListResponse{}, nil
	}
	var friends []string
	for _, relationship := range relationships {
		friends = append(friends, relationship.Target)
	}
	return models.CommonFriendListResponse{Success: true, Friends: friends, Count: len(friends)}, nil

}

func (svc *service) SubscribeFromEmail(request models.SubscribeRequest) (models.SubscribeResponse, error) {
	relationships, err := svc.repository.SubscribeFromEmail(request)
	if err != nil {
		return models.SubscribeResponse{}, err
	}
	if (relationships != models.Relationship{}) {
		return models.SubscribeResponse{Success: true}, nil
	}
	return models.SubscribeResponse{Success: false}, nil
}

func (svc *service) BlockSubscribeByEmail(request models.BlockSubscribeRequest) (models.BlockSubscribeResponse, error) {
	relationship, err := svc.repository.BlockSubscribeByEmail(request)
	if err != nil {
		return models.BlockSubscribeResponse{}, err
	}
	if relationship.FriendBlocked == true {
		return models.BlockSubscribeResponse{Success: true}, nil
	}
	return models.BlockSubscribeResponse{Success: false}, nil
}

func (svc *service) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) (models.GetSubscribingEmailListResponse, error) {
	response := models.GetSubscribingEmailListResponse{Success: false}
	if request == (models.GetSubscribingEmailListRequest{}) {
		return models.GetSubscribingEmailListResponse{}, errors.New("Invalid Request")
	}
	relationship, err := svc.repository.GetSubscribingEmailListByEmail(request)
	if err != nil {
		return models.GetSubscribingEmailListResponse{}, err
	} else {
		response.Success = true
		var recipients []string
		for _, re := range relationship {
			recipients = append(recipients, re.Target)
		}
		response.Recipients = recipients
	}
	return response, nil
}
