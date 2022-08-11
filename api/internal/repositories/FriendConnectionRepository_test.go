package repositories

import (
	"testing"

	"golang_project/api/cmd/serverd/database"
	"golang_project/api/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	repo FriendConnectionRepository = New(database.GetInstance())
)

// 1.
func TestFindFriendsByEmailWithSuccessfulCase(t *testing.T) {
	result, _ := repo.FindFriendsByEmail("thehaohcm@yahoo.com.vn")
	expectedResult := []string{
		"hao.nguyen@s3corp.com.vn",
	}
	assert.Equal(t, expectedResult, result)
}

func TestFindFriendsByEmailWithNoResult(t *testing.T) {
	result, _ := repo.FindFriendsByEmail("test@test.com")
	expectedResult := []string(nil)
	assert.Equal(t, expectedResult, result)
}

func TestFindFriendsByEmailWithEmptyRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.FindFriendsByEmail("")
}

func TestFindFriendsByEmailWithInvalidEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.FindFriendsByEmail("abc")
}

// 2.
func TestFindCommonFriendsByEmailsWithSuccessfulCase(t *testing.T) {
	result, _ := repo.FindCommonFriendsByEmails([]string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"})
	expectedRs := []string{"chinh.nguyen@s3corp.com.vn"}
	assert.Equal(t, expectedRs, result)
}

func TestFindCommonFriendsByEmailsWithEmptyResponse(t *testing.T) {
	result, _ := repo.FindCommonFriendsByEmails([]string{"thehaohcm@yahoo.com.vn", "hung.tong@s3corp.com.vn"})
	expectedRs := []string(nil)
	assert.Equal(t, expectedRs, result)
}

func TestFindCommonFriendsByEmailsWithNilRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.FindCommonFriendsByEmails(nil)
}

func TestFindCommonFriendsByEmailsWithEmptyEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.FindCommonFriendsByEmails([]string{""})
}

func TestFindCommonFriendsByEmailsWithInvalidEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	repo.FindCommonFriendsByEmails([]string{"test"})
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
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.CreateFriendConnection([]string{
		"abc@def.com",
		"abc1@def.com",
	})

	assert.Equal(t, true, result)
}

func TestCreateFriendConnectionWithNilRequest(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	mockRepo.CreateFriendConnection(nil)
}

func TestCreateFriendConnectionWithEmptyRequest(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.CreateFriendConnection([]string{""})
}

func TestCreateFriendConnectionWithInvalidEmail(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockRepo.CreateFriendConnection([]string{"test"})
}

func TestCreateFriendConnectionWithExceesEmails(t *testing.T) {

	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	mockRepo.CreateFriendConnection([]string{"hao.nguyen@s3corp.com.vn", "thehaohcm@yahoo.com.vn", "thehaohcm@gmai.com"})
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
	sqlMock.ExpectExec("INSERT INTO public.subscribers").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "chinh.nguyen@s3corp.com.vn"})
	assert.Equal(t, true, result)
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
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "thehaohcm@gmail.com"})
	assert.Equal(t, true, result)
}

func TestBlockSubscribeByEmailWithSuccessfulCaseAndHaveFriend(t *testing.T) {
	var mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()

	var mockRepo FriendConnectionRepository = New(mockDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO public.friends").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	result, _ := mockRepo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "chinh.nguyen@s3corp.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	assert.Equal(t, true, result)
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

// 6.
func TestGetSubscribingEmailListByEmailWithSuccessfulCaseAndEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world, kate@example.com"})
	expectedRs := []string{
		"hao.nguyen@s3corp.com.vn",
		"abc@gmail.com",
		"kate@example.com",
	}
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulCaseNotEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world"})
	expectedRs := []string{
		"hao.nguyen@s3corp.com.vn", "abc@gmail.com",
	}
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulAndEmptyResponse(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "hello world"})
	expectedRs := []string(nil)
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
