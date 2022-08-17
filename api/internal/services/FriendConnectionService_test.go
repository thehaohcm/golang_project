package services

import (
	"testing"

	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// external
func TestCreateUserSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result := myService.CreateUser(models.CreatingUserRequest{Email: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.CreatingUserResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestCreateUserInvalidEmailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()
	myService.CreateUser(models.CreatingUserRequest{Email: "hao.nguyen"})
}

func TestCreateUserNilCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()
	myService.CreateUser(models.CreatingUserRequest{})
}

// 1.
func TestFriendConnectionSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}})
	expectedRs := models.FriendConnectionResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestFriendConnectionFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result := myService.CreateConnection(models.FriendConnectionRequest{Friends: []string{}})
	expectedRs := models.FriendConnectionResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

// 2.
func TestShowFriendsByEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.FriendListRequest{
		Email: "thehaohcm@yahoo.com.vn",
	}

	response := myService.GetFriendConnection(request)

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
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function doesn't panic")
		}
	}()
	myService.GetFriendConnection(models.FriendListRequest{})

}

func TestShowFriendsByEmailWithEmptyResponse(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.FriendListRequest{
		Email: "test@test.com",
	}

	response := myService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: false,
		Friends: []string(nil),
		Count:   0,
	}
	assert.Equal(t, exp, response)
}

// 3.
func TestShowCommonFriendListSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	request := models.CommonFriendListRequest{
		Friends: []string{"thehaohcm@yahoo.com.vn", "chinh.nguyen@s3corp.com.vn"},
	}

	response := myService.ShowCommonFriendList(request)

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
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()
	myService.ShowCommonFriendList(models.CommonFriendListRequest{Friends: []string{"hao.nguyen"}})
}

func TestShowCommonFriendListEmptyModel(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response := myService.ShowCommonFriendList(models.CommonFriendListRequest{})

	exp := models.CommonFriendListResponse{}
	assert.Equal(t, exp, response)
}

// 4.
func TestSubscribeFromEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.SubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailFailCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailWithEmptyRequestor(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(nil).Once()
	myService := New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

func TestSubscribeFromEmailWithEmptyTarget(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	// repoMock.On("SubscribeFromEmail", test.AnythingOfType("models.SubscribeRequest")).Return(nil).Once()
	myService := New(repoMock)
	result := myService.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
	expectedRs := models.SubscribeResponse{Success: false}
	assert.Equal(t, expectedRs, result)
}

// 5.
func TestBlockSubscribeByEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	// repoMock.On("BlockSubscribeByEmail", test.AnythingOfType("models.BlockSubscribeRequest")).Return(true, nil).Once()
	myService := New(repoMock)
	result := myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	expectedRs := models.BlockSubscribeResponse{Success: true}
	assert.Equal(t, expectedRs, result)
}

func TestBlockSubscribeByEmailFailCase(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()

	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{})
}

func TestBlockSubscribeByEmailWithEmptyTarget(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()

	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn"})
}

func TestBlockSubscribeByEmailWithEmptyRequestor(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()

	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)
	myService.BlockSubscribeByEmail(models.BlockSubscribeRequest{Target: "thehaohcm@yahoo.com.vn"})
}

// 6.
func TestGetSubscribingEmailListWithEmailSuccessfulCase(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	model := models.GetSubscribingEmailListRequest{
		Sender: "thehaohcm@yahoo.com.vn",
		Text:   "helloworld! kate@example.com",
	}

	response := myService.GetSubscribingEmailListByEmail(model)

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
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	model := models.GetSubscribingEmailListRequest{
		Sender: "dfa@yahoo.com.vn",
		Text:   "helloworld!",
	}

	response := myService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success:    true,
		Recipients: nil,
	}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListWithEmailEmptyModel(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{})

	exp := models.GetSubscribingEmailListResponse{}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListWithInvalidEmail(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()
	myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm", Text: "abc"})
}

func TestGetSubscribingEmailListWithNilSender(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the function is not panic")
		}
	}()
	myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Text: "abc"})
}

func TestGetSubscribingEmailListWithEmptyReponse(t *testing.T) {
	repoMock := &FriendConnectionRepoMock{}
	myService := New(repoMock)

	response := myService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "abc"})
	expRs := models.GetSubscribingEmailListResponse{Success: true, Recipients: nil}
	assert.Equal(t, expRs, response)
}

type FriendConnectionRepoMock struct {
	mock.Mock
}

func (f *FriendConnectionRepoMock) CreateUser(request models.CreatingUserRequest) (models.User, error) {
	if valid, err := pkg.CheckValidEmail(request.Email); !valid || err != nil {
		return models.User{}, err
	}
	return models.User{Email: request.Email}, nil
}

func (f *FriendConnectionRepoMock) FindFriendsByEmail(request models.FriendListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmail(request.Email); !valid || err != nil {
		panic("invalid email address")
	}

	if request.Email == "thehaohcm@yahoo.com.vn" {
		return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}}, nil
	}
	return []models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) FindCommonFriendsByEmails(request models.CommonFriendListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails(request.Friends); !valid || err != nil {
		panic("invalid email address")
	}
	for _, item := range request.Friends {
		if item == "thehaohcm@yahoo.com.vn" {
			return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}}, nil
		}
	}

	return []models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) CreateFriendConnection(request models.FriendConnectionRequest) (models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails(request.Friends); !valid || err != nil {
		return models.Relationship{}, nil
	}
	if len(request.Friends) > 0 {
		return models.Relationship{Is_friend: true}, nil
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
	if valid, err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); !valid || err != nil {
		panic("invalid email address")
	}
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return models.Relationship{Target: "hao.nguyen@s3corp.com.vn", Friend_blocked: true}, nil
	}
	return models.Relationship{}, nil
}

func (f *FriendConnectionRepoMock) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmail(req.Sender); !valid || err != nil {
		panic("invalid email address")
	}
	if req.Sender == "thehaohcm@yahoo.com.vn" && req.Text == "helloworld! kate@example.com" {
		return []models.Relationship{{Target: "hao.nguyen@s3corp.com.vn"}, {Target: "kate@example.com"}}, nil
	}
	return []models.Relationship{}, nil
}
