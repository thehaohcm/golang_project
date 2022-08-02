package test

import (
	"golang_project/models"
)

type FriendConnectionRepository interface {
	FindFriendsByEmail(email string) []string
	FindCommonFriendsByEmails(emails []string) []string
	CreateFriendConnection(emails []string) bool
	SubscribeFromEmail(req models.SubscribeRequest) bool
	BlockSubscribeByEmail(req models.BlockSubscribeRequest) bool
	GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse
}

type repository struct {
	// db *sql.DB
}

func New() FriendConnectionRepository {
	return &repository{
		// db: utils.GetInstance(),
	}
}

//1.
func (repo *repository) CreateFriendConnection(emails []string) bool {
	return true
}

//2.
func (repo *repository) FindFriendsByEmail(email string) []string {
	var friends = []string{"hao.nguyen@s3corp.com.vn", "thehaohcm@yahoo.com.vn"}
	return friends
}

//3.
func (repo *repository) FindCommonFriendsByEmails(emails []string) []string {
	var friends = []string{"hao.nguyen@s3corp.com.vn", "thehaohcm@yahoo.com.vn"}
	return friends
}

//4.
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) bool {
	return true
}

//5.
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) bool {
	return true
}

func (repo *repository) hasFriendConnection(requestor string, target string) bool {
	return true
}

//6.
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	var res models.GetSubscribingEmailListResponse
	res.Recipients = []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}
	res.Success = true
	return res
}
