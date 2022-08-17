package repositories

import (
	"errors"
	"testing"

	"golang_project/api/cmd/golang_project/database"
	"golang_project/api/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	repo FriendConnectionRepository = New(database.GetInstance())
)

// external
func TestCreateUserWithSuccessfulCase(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.user_account").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateUser(models.CreatingUserRequest{"abc@def.com"})
	expectedResult := models.User{Email: "abc@def.com"}
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.user_account").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateUser(models.CreatingUserRequest{"abc"})
	expectedResult := models.User{}
	errExpected := errors.New("invalid email address")
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, errExpected, err)
}

func TestCreateUserWithEmptyBody(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.user_account").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateUser(models.CreatingUserRequest{""})
	expectedResult := models.User{}
	errExpected := errors.New("invalid email address")
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, errExpected, err)
}

// 1. TODO: cannot run successfully
func TestFindFriendsByEmailWithSuccessfulCase(t *testing.T) {
	result, err := repo.FindFriendsByEmail(models.FriendListRequest{"thehaohcm@yahoo.com.vn"})
	expectedResult := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "thehaohcm@yahoo.com.vn", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "thehaohcm@yahoo.com.vn", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
}

func TestFindFriendsByEmailWithNoResult(t *testing.T) {
	result, err := repo.FindFriendsByEmail(models.FriendListRequest{"test@test.com"})
	expectedResult := []models.Relationship([]models.Relationship(nil))
	assert.Equal(t, expectedResult, result)
	assert.IsType(t, nil, err)
}

func TestFindFriendsByEmailWithEmptyRequest(t *testing.T) {
	result, err := repo.FindFriendsByEmail(models.FriendListRequest{""})
	expectedResult := []models.Relationship{}
	assert.Equal(t, expectedResult, result)
	assert.Error(t, err, errors.New("email address is emtpy"))
}

func TestFindFriendsByEmailWithInvalidEmailRequest(t *testing.T) {
	result, err := repo.FindFriendsByEmail(models.FriendListRequest{"abc"})
	expectedResult := []models.Relationship{}
	assert.Equal(t, expectedResult, result)
	assert.Error(t, err, errors.New("email address is emtpy"))
}

// 2. TODO: cannot run succesfully
func TestFindCommonFriendsByEmailsWithSuccessfulCase(t *testing.T) {
	result, err := repo.FindCommonFriendsByEmails(models.CommonFriendListRequest{[]string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}})
	expectedRs := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedRs, result)
	assert.IsType(t, nil, err)
}

func TestFindCommonFriendsByEmailsWithEmptyResponse(t *testing.T) {
	result, err := repo.FindCommonFriendsByEmails(models.CommonFriendListRequest{[]string{"thehaohcm@yahoo.com.vn", "hung.tong@s3corp.com.vn"}})
	expectedRs := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedRs, result)
	assert.IsType(t, nil, err)
}

func TestFindCommonFriendsByEmailsWithNilRequest(t *testing.T) {
	result, err := repo.FindCommonFriendsByEmails(models.CommonFriendListRequest{})
	assert.Equal(t, []models.Relationship{}, result)
	assert.Error(t, err, errors.New("email address is empty"))
}

func TestFindCommonFriendsByEmailsWithEmptyEmailRequest(t *testing.T) {
	result, err := repo.FindCommonFriendsByEmails(models.CommonFriendListRequest{})
	assert.Equal(t, []models.Relationship{}, result)
	assert.Error(t, err, errors.New("email address is emtpy"))
}

func TestFindCommonFriendsByEmailsWithInvalidEmailRequest(t *testing.T) {
	result, err := repo.FindCommonFriendsByEmails(models.CommonFriendListRequest{Friends: []string{"test"}})
	assert.Equal(t, []models.Relationship{}, result)
	assert.Error(t, err, errors.New("invalid email address"))
}

// 3.
func TestCreateFriendConnectionWithSuccessfulCase(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateFriendConnection(models.FriendConnectionRequest{[]string{"abc@def.com", "abc1@def.com"}})
	expectedResult := models.Relationship(models.Relationship{Requestor: "abc@def.com", Target: "abc1@def.com", Is_friend: true, Friend_blocked: false, Subscribed: false, Subscribe_block: false})
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
}

func TestCreateFriendConnectionWithNilRequest(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateFriendConnection(models.FriendConnectionRequest{})
	assert.Equal(t, models.Relationship{}, result)
	assert.Error(t, err, errors.New("invalid request"))
}

func TestCreateFriendConnectionWithEmptyRequest(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateFriendConnection(models.FriendConnectionRequest{})

	assert.Equal(t, models.Relationship{}, result)
	assert.Error(t, err, errors.New("email address is emtpy"))
}

func TestCreateFriendConnectionWithInvalidEmail(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateFriendConnection(models.FriendConnectionRequest{[]string{"test"}})
	assert.Equal(t, models.Relationship{}, result)
	assert.Error(t, err, errors.New("invalid email address"))
}

func TestCreateFriendConnectionWithExceesEmails(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.CreateFriendConnection(models.FriendConnectionRequest{[]string{"hao.nguyen@s3corp.com.vn", "thehaohcm@yahoo.com.vn", "thehaohcm@gmail.com"}})
	assert.Equal(t, models.Relationship{}, result)
	assert.Error(t, err, errors.New("invalid request"))
}

// 4.
func TestSubscribeFromEmailWithSuccessfulCase(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, err := mockRepo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "chinh.nguyen@s3corp.com.vn"})
	expectedResult := models.Relationship(models.Relationship{Requestor: "thehaohcm@yahoo.com.vn", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: true, Subscribe_block: false})
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, nil, err)
}

func TestSubscribeFromEmailWithInvalidEmail(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen@s3corp.com.vn"})
}

func TestSubscribeFromEmailWithInvalidEmails(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen"})
}

func TestSubscribeFromEmailWithNilReq(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.SubscribeFromEmail(models.SubscribeRequest{})
}

// 5.
func TestBlockSubscribeByEmailWithSuccessfulCaseAndHaveNoFriend(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "thehaohcm@gmail.com"})
	expectedResult := models.Relationship(models.Relationship{Requestor: "thehaohcm@yahoo.com.vn", Target: "thehaohcm@gmail.com", Is_friend: false, Friend_blocked: true, Subscribed: false, Subscribe_block: false})
	assert.Equal(t, expectedResult, result)
}

func TestBlockSubscribeByEmailWithSuccessfulCaseAndHaveFriend(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.relationship").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "chinh.nguyen@s3corp.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	expectedResult := models.Relationship(models.Relationship{Requestor: "chinh.nguyen@s3corp.com.vn", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: true, Subscribed: false, Subscribe_block: false})
	assert.Equal(t, expectedResult, result)
}

func TestBlockSubscribeByEmailInvalidEmails(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen"})
}

func TestBlockSubscribeByEmailWithNilRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{})
}

func TestBlockSubscribeByEmailWithNilRequestor(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Target: "chinh.nguyen"})
}

func TestBlockSubscribeByEmailWithNilTarget(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm"})
}

// 6. TODO: cannot run successfully
func TestGetSubscribingEmailListByEmailWithSuccessfulCaseAndEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world, kate@example.com"})
	expectedRs := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "", Target: "kate@example.com", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "thehaohcm@yahoo.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulCaseNotEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world"})
	expectedRs := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "", Target: "thehaohcm@yahoo.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulAndEmptyResponse(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "hello world"})
	expectedRs := []models.Relationship([]models.Relationship{models.Relationship{Requestor: "", Target: "chinh.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "hung.tong@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}, models.Relationship{Requestor: "", Target: "hao.nguyen@s3corp.com.vn", Is_friend: false, Friend_blocked: false, Subscribed: false, Subscribe_block: false}})
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithNilSender(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Text: "hello world"})
}

func TestGetSubscribingEmailListByEmailWithInvalidEmail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm", Text: "hello world"})
}
