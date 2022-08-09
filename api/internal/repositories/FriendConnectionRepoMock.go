package repositories

import (
	"golang_project/api/internal/models"

	"github.com/stretchr/testify/mock"
)

type FriendConnectionRepoMock struct {
	mock.Mock
}

func (f *FriendConnectionRepoMock) FindFriendsByEmail(email string) ([]string, error) {
	return []string{}, nil
}
func (f *FriendConnectionRepoMock) FindCommonFriendsByEmails(emails []string) ([]string, error) {
	return []string{}, nil
}
func (f *FriendConnectionRepoMock) CreateFriendConnection(emails []string) (bool, error) {
	if len(emails) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) SubscribeFromEmail(req models.SubscribeRequest) (bool, error) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, error) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]string, error) {
	return []string{}, nil
}
