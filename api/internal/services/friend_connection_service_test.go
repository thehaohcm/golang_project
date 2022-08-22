package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"
)

func TestCreateUserSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.CreateUser(models.CreatingUserRequest{Email: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.CreatingUserResponse{Success: true}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestCreateUserInvalidEmailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.CreateUser(models.CreatingUserRequest{Email: "hao.nguyen"})
	assert.Equal(t, models.CreatingUserResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestCreateUserNilCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.CreateUser(models.CreatingUserRequest{})
	assert.Equal(t, models.CreatingUserResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestFriendConnectionSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}})
	expectedRs := models.FriendConnectionResponse{Success: true}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestFriendConnectionFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{}})
	expectedRs := models.FriendConnectionResponse{Success: false}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, errors.New("Email address is empty"), err)
}

func TestShowFriendsByEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.FriendListRequest{
		Email: "thehaohcm@yahoo.com.vn",
	}

	response, err := myService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
	assert.Equal(t, nil, err)
}

func TestShowFriendsByEmailEmptyModel(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.GetFriendConnection(models.FriendListRequest{})
	assert.Equal(t, models.FriendListResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestShowFriendsByEmailWithEmptyResponse(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.FriendListRequest{
		Email: "test@test.com",
	}

	response, err := myService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: false,
		Friends: []string(nil),
		Count:   0,
	}
	assert.Equal(t, exp, response)
	assert.Equal(t, nil, err)
}

func TestShowCommonFriendListSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.CommonFriendListRequest{
		Friends: []string{"thehaohcm@yahoo.com.vn", "chinh.nguyen@s3corp.com.vn"},
	}

	response, err := myService.ShowCommonFriendList(request)

	exp := models.CommonFriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
	assert.Equal(t, nil, err)
}

func TestShowCommonFriendListWithInvalidEmail(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.ShowCommonFriendList(models.CommonFriendListRequest{Friends: []string{"hao.nguyen"}})
	assert.Equal(t, models.CommonFriendListResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestShowCommonFriendListEmptyModel(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response, err := myService.ShowCommonFriendList(models.CommonFriendListRequest{})

	exp := models.CommonFriendListResponse{}
	assert.Equal(t, exp, response)
	assert.Equal(t, errors.New("Email address is empty"), err)
}

func TestSubscribeFromEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.SubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestSubscribeFromEmailFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.SubscribeFromEmail(models.SubscribeRequest{})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestSubscribeFromEmailWithEmptyRequestor(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.SubscribeFromEmail(models.SubscribeRequest{Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestSubscribeFromEmailWithEmptyTarget(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestBlockSubscribeByEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.BlockSubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
	assert.Equal(t, nil, err)
}

func TestBlockSubscribeByEmailFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{})
	assert.Equal(t, models.BlockSubscribeResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestBlockSubscribeByEmailWithEmptyTarget(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
	assert.Equal(t, models.BlockSubscribeResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestBlockSubscribeByEmailWithEmptyRequestor(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Target: "thehaohcm@yahoo.com.vn"})
	assert.Equal(t, models.BlockSubscribeResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestGetSubscribingEmailListWithEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	model := models.GetSubscribingEmailListRequest{
		Sender: "thehaohcm@yahoo.com.vn",
		Text:   "helloworld! kate@example.com",
	}

	response, err := myService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success: true,
		Recipients: []string{
			"hao.nguyen@s3corp.com.vn",
			"kate@example.com",
		},
	}
	assert.Equal(t, exp, response)
	assert.Equal(t, nil, err)
}

func TestGetSubscribingEmailListWithEmailFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	model := models.GetSubscribingEmailListRequest{
		Sender: "dfa@yahoo.com.vn",
		Text:   "helloworld!",
	}

	response, err := myService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success:    true,
		Recipients: nil,
	}
	assert.Equal(t, exp, response)
	assert.Equal(t, nil, err)
}

func TestGetSubscribingEmailListWithEmailEmptyModel(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response, err := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{})

	exp := models.GetSubscribingEmailListResponse{}
	assert.Equal(t, exp, response)
	assert.Equal(t, errors.New("Invalid Request"), err)
}

func TestGetSubscribingEmailListWithInvalidEmail(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm", Text: "abc"})
	assert.Equal(t, models.GetSubscribingEmailListResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestGetSubscribingEmailListWithNilSender(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result, err := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Text: "abc"})
	assert.Equal(t, models.GetSubscribingEmailListResponse{}, result)
	assert.IsType(t, errors.New(""), err)
}

func TestGetSubscribingEmailListWithEmptyReponse(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response, err := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "abc"})
	expRs := models.GetSubscribingEmailListResponse{Success: true, Recipients: nil}
	assert.Equal(t, expRs, response)
	assert.Equal(t, nil, err)
}

type FriendConnectionRepoMock struct {
	mock.Mock
}

func (f *FriendConnectionRepoMock) CreateUser(request models.CreatingUserRequest) (models.User, error) {
	if err := pkg.CheckValidEmail(request.Email); err != nil {
		return models.User{}, err
	}
	return models.User{Email: request.Email}, nil
}

func (f *FriendConnectionRepoMock) FindFriendsByEmail(request models.FriendListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmail(request.Email); err != nil {
		return []models.Relationship{}, err
	}

	if request.Email == "thehaohcm@yahoo.com.vn" {
		return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}}, nil
	}
	return []models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) FindCommonFriendsByEmails(request models.CommonFriendListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return []models.Relationship{}, err
	}
	for _, item := range request.Friends {
		if item == "thehaohcm@yahoo.com.vn" {
			return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}}, nil
		}
	}

	return []models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) CreateFriendConnection(request models.FriendConnectionRequest) (models.Relationship, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return models.Relationship{}, err
	}
	if len(request.Friends) > 0 {
		return models.Relationship{IsFriend: true}, nil
	}
	return models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) SubscribeFromEmail(req models.SubscribeRequest) (models.Relationship, error) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return models.Relationship{Target: "hao.nguyen@s3corp.com.vn"}, nil
	}
	return models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (models.Relationship, error) {
	if err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); err != nil {
		return models.Relationship{}, err
	}
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return models.Relationship{Target: "hao.nguyen@s3corp.com.vn", FriendBlocked: true}, nil
	}
	return models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmail(req.Sender); err != nil {
		return []models.Relationship{}, err
	}
	if req.Sender == "thehaohcm@yahoo.com.vn" && req.Text == "helloworld! kate@example.com" {
		return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}, {Target: "kate@example.com"}}, nil
	}
	return []models.Relationship{}, nil
}
