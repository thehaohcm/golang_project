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

// New function used for initializing a service
// pass a FriendConnectionRepository as parameter
func New(repo repositories.FriendConnectionRepository) FriendConnectionService {
	return &service{
		repository: repo,
	}
}

// CreateUse function works as a service function for creating an new user
// pass a CreatingUserRequest model as parameter
// return a CreatingUserResponse model and an error type
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

// CreateConnection function works as a service function for creating friend connection between 2 user emails
// pass a FriendConnectionRequest model as parameter
// return a FriendConnectionResponse model and an error type
func (svc *service) CreateConnection(request models.FriendConnectionRequest) (models.FriendConnectionResponse, error) {
	_, err := svc.repository.CreateFriendConnection(request)
	if err != nil {
		return models.FriendConnectionResponse{}, err
	}
	return models.FriendConnectionResponse{Success: true}, nil
}

// GetFriendConnection function works as a service function for getting a friend list by an email address
// pass a FriendListRequest model as parameter
// return a FriendListResponse model and an error type
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

// ShowCommonFriendList function works as a service function for getting a list of common friends between two email addresses
// pass a CommonFriendListRequest model as parameter
// return a CommonFriendListResponse model and an error type
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

// SubscribeFromEmail function works as a service function for creating a subscribe from an email address to another one
// pass SubscribeRequest model as parameter
// return a SubscribeResponse model and an error type
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

// BlockSubscribeByEmail function works as a service function for creating a block subscribe update from an email address to another one
// pass a BlockSubscribeRequest model as parameter
// return a BlockSubscribeResponse model and an error type
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

// GetSubscribingEmailListByEmail function works as a service function for getting a list of subscribe email by an email address
// pass a GetSubscribingEmailListRequest model as parameter
// return a GetSubscribingEmailListRequest model and an error type
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
