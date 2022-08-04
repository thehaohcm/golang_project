package test

import (
	"golang_project/controllers"
	"golang_project/models"
	"golang_project/repositories"
	"golang_project/services"
	test "golang_project/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	friendConnectionRepository repositories.FriendConnectionRepository = repositories.New()
	friendConnectionService    services.FriendConnectionService        = services.New(friendConnectionRepository)
	friendConnectionController controllers.FriendConnectionController  = controllers.New(friendConnectionService)
)

//1.
func TestFriendConnectionSuccessfulCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("CreateConnection", test.AnythingOfType("*map[string]interface{}")).Return(true, nil)
	myService := services.New(repoMock)
	result := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}})
	expectedRs := models.FriendConnectionResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestFriendConnectionFailCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("CreateConnection", test.AnythingOfType("*map[string]interface{}")).Return(false, nil)
	myService := services.New(repoMock)
	result := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{}})
	expectedRs := models.FriendConnectionResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

//2.
func TestShowFriendsByEmailSuccessfulCase(t *testing.T) {

	request := models.FriendListRequest{
		Email: "thehaohcm@yahoo.com.vn",
	}

	response := friendConnectionService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
}

func TestShowFriendsByEmailEmptyModel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	friendConnectionService.GetFriendConnection(models.FriendListRequest{})
}

func TestShowFriendsByEmailWithEmptyResponse(t *testing.T) {

	request := models.FriendListRequest{
		Email: "test@test.com",
	}

	response := friendConnectionService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: false,
		Friends: []string(nil),
		Count:   0,
	}
	assert.Equal(t, exp, response)
}

//3.
func TestShowCommonFriendListSuccessfulCase(t *testing.T) {

	request := models.CommonFriendListRequest{
		Friends: []string{"thehaohcm@yahoo.com.vn", "chinh.nguyen@s3corp.com.vn"},
	}

	response := friendConnectionService.ShowCommonFriendList(request)

	exp := models.CommonFriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
}

func TestShowCommonFriendListWithInvalidEmail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	friendConnectionService.ShowCommonFriendList(models.CommonFriendListRequest{Friends: []string{"hao.nguyen"}})
}

func TestShowCommonFriendListEmptyModel(t *testing.T) {

	response := friendConnectionService.ShowCommonFriendList(models.CommonFriendListRequest{})

	exp := models.CommonFriendListResponse{}
	assert.Equal(t, exp, response)
}

//4.
func TestSubscribeFromEmailSuccessfulCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(true, nil).Once()
	myService := services.New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com .vn"})
	expectedRs := models.SubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailFailCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(nil).Once()
	myService := services.New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailWithEmptyRequestor(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(nil).Once()
	myService := services.New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailWithEmptyTarget(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(nil).Once()
	myService := services.New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

//5.
func TestBlockSubscribeByEmailSuccessfulCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("BlockSubscribeByEmail", test.AnythingOfType("models.BlockSubscribeRequest")).Return(true, nil).Once()
	myService := services.New(repoMock)
	result := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yoo.com.vn", Target: "hao.nguyen@s3corp.com .vn"})
	expectedRs := models.BlockSubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}
func TestBlockSubscribeByEmailFailCase(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("BlockSubscribeByEmail", test.AnythingOfType("models.BlockSubscribeRequest")).Return(false, nil).Once()
	myService := services.New(repoMock)
	result := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{})
	expectedRs := models.BlockSubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestBlockSubscribeByEmailWithEmptyTarget(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("BlockSubscribeByEmail", test.AnythingOfType("models.BlockSubscribeRequest")).Return(false, nil).Once()
	myService := services.New(repoMock)
	result := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
	expectedRs := models.BlockSubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestBlockSubscribeByEmailWithEmptyRequestor(t *testing.T) {
	repoMock := &test.FriendConnectionRepoMock{}
	// repoMock.On("BlockSubscribeByEmail", test.AnythingOfType("models.BlockSubscribeRequest")).Return(false, nil).Once()
	myService := services.New(repoMock)
	result := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Target: "thehaohcm@yahoo.com.vn"})
	expectedRs := models.BlockSubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

//6.
func TestGetSubscribingEmailListWithEmailSuccessfulCase(t *testing.T) {

	model := models.GetSubscribingEmailListRequest{
		Sender: "thehaohcm@yahoo.com.vn",
		Text:   "helloworld! kate@example.com",
	}

	response := friendConnectionService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success: true,
		Recipients: []string{
			"hao.nguyen@s3corp.com.vn",
			"kate@example.com",
		},
	}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListWithEmailFailCase(t *testing.T) {

	model := models.GetSubscribingEmailListRequest{
		Sender: "dfa@yahoo.com.vn",
		Text:   "helloworld!",
	}

	response := friendConnectionService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success:    false,
		Recipients: nil,
	}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListWithEmailEmptyModel(t *testing.T) {

	response := friendConnectionService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{})

	exp := models.GetSubscribingEmailListResponse{}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListWithInvalidEmail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	friendConnectionService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm", Text: "abc"})
}

func TestGetSubscribingEmailListWithNilSender(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	friendConnectionService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Text: "abc"})
}

func TestGetSubscribingEmailListWithEmptyReponse(t *testing.T) {
	response := friendConnectionService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "abc"})
	expRs := models.GetSubscribingEmailListResponse{Success: false, Recipients: nil}
	assert.Equal(t, expRs, response)
}
