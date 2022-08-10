package repositories

import (
	"database/sql"
	"testing"

	"golang_project/api/internal/models"

	"github.com/stretchr/testify/assert"
)

var (
	repo FriendConnectionRepository = New()
)

//1.
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
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.FindFriendsByEmail("")
}

func TestFindFriendsByEmailWithInvalidEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.FindFriendsByEmail("abc")
}

//2.
func TestFindCommonFriendsByEmailsWithSuccessfulCase(t *testing.T) {
	result, _ := repo.FindCommonFriendsByEmails([]string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"})
	expectedRs := []string{"chinh.nguyen@s3corp.com.vn"}
	assert.Equal(t, expectedRs, result)
}

func TestFindCommonFriendsByEmailsWithEmptyResponse(t *testing.T) {
	result, _ := repo.FindCommonFriendsByEmails([]string{"thehaohcm@yahoo.com.vn", "thehaohcm@gmail.com"})
	expectedRs := []string(nil)
	assert.Equal(t, expectedRs, result)
}

func TestFindCommonFriendsByEmailsWithNilRequest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "empty request", msg)
		}
	}()
	repo.FindCommonFriendsByEmails(nil)
}

func TestFindCommonFriendsByEmailsWithEmptyEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.FindCommonFriendsByEmails([]string{""})
}

func TestFindCommonFriendsByEmailsWithInvalidEmailRequest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.FindCommonFriendsByEmails([]string{"test"})
}

//3.
func TestCreateFriendConnectionWithSuccessfulCase(t *testing.T) {

	result, _ := repo.CreateFriendConnection([]string{
		"abc@def.com",
		"abc1@def.com",
	})

	//rollback db
	rollbackCtx(tx)

	assert.Equal(t, true, result)
}

func TestCreateFriendConnectionWithNilRequest(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()

	repo.CreateFriendConnection(nil)
}

func TestCreateFriendConnectionWithEmptyRequest(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.CreateFriendConnection([]string{""})
}

func TestCreateFriendConnectionWithInvalidEmail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.CreateFriendConnection([]string{"test"})
}

func TestCreateFriendConnectionWithExceesEmails(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()

	repo.CreateFriendConnection([]string{"hao.nguyen@s3corp.com.vn", "thehaohcm@yahoo.com.vn", "thehaohcm@gmai.com"})
}

//4.
func TestSubscribeFromEmailWithSuccessfulCase(t *testing.T) {
	result, _ := repo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "chinh.nguyen@s3corp.com.vn"})
	assert.Equal(t, true, result)
}

func TestSubscribeFromEmailWithInvalidEmail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen@s3corp.com.vn"})
}

func TestSubscribeFromEmailWithInvalidEmails(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.SubscribeFromEmail(models.SubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen"})
}

func TestSubscribeFromEmailWithNilReq(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.SubscribeFromEmail(models.SubscribeRequest{})
}

//5.
func TestBlockSubscribeByEmailWithSuccessfulCaseAndHaveNoFriend(t *testing.T) {
	result, _ := repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm@yahoo.com.vn", Target: "thehaohcm@gmail.com"})
	assert.Equal(t, true, result)
}

func TestBlockSubscribeByEmailWithSuccessfulCaseAndHaveFriend(t *testing.T) {
	result, _ := repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "chinh.nguyen@s3corp.com.vn", Target: "hao.nguyen@s3corp.com.vn"})
	assert.Equal(t, true, result)
}

func TestBlockSubscribeByEmailInvalidEmails(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm", Target: "chinh.nguyen"})
}

func TestBlockSubscribeByEmailWithNilRequest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{})
}

func TestBlockSubscribeByEmailWithNilRequestor(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Target: "chinh.nguyen"})
}

func TestBlockSubscribeByEmailWithNilTarget(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.BlockSubscribeByEmail(models.BlockSubscribeRequest{Requestor: "thehaohcm"})
}

//6.
func TestGetSubscribingEmailListByEmailWithSuccessfulCaseAndEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world, kate@example.com"})
	expectedRs := []string{
		"hao.nguyen@s3corp.com.vn",
		"kate@example.com",
	}
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulCaseNotEmailInText(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm@yahoo.com.vn", Text: "hello world"})
	expectedRs := []string{
		"hao.nguyen@s3corp.com.vn",
	}
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithSuccessfulAndEmptyResponse(t *testing.T) {
	result, _ := repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "hung.tong@s3corp.com.vn", Text: "hello world"})
	expectedRs := []string{}
	assert.Equal(t, expectedRs, result)
}

func TestGetSubscribingEmailListByEmailWithNilSender(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Text: "hello world"})
}

func TestGetSubscribingEmailListByEmailWithInvalidEmail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(string)
			assert.Equal(t, "invalid request", msg)
		}
	}()
	repo.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{Sender: "thehaohcm", Text: "hello world"})
}

func rollbackCtx(tx *sql.Tx) {
	if tx != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err)
		}
	}
}
