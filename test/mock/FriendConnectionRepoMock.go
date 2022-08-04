package test

import (
	"database/sql"
	"golang_project/models"

	"github.com/stretchr/testify/mock"
)

type FriendConnectionRepoMock struct {
	mock.Mock
}

func (f *FriendConnectionRepoMock) FindFriendsByEmail(email string) []string {
	return []string{}
}
func (f *FriendConnectionRepoMock) FindCommonFriendsByEmails(emails []string) []string {
	return []string{}
}
func (f *FriendConnectionRepoMock) CreateFriendConnection(emails []string) (bool, *sql.Tx) {
	if len(emails) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) SubscribeFromEmail(req models.SubscribeRequest) (bool, *sql.Tx) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, *sql.Tx) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return true, nil
	}
	return false, nil
}
func (f *FriendConnectionRepoMock) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	return models.GetSubscribingEmailListResponse{}
}
