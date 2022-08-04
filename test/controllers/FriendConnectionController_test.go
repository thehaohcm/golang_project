package test

import (
	"encoding/json"
	"golang_project/models"
	"golang_project/routes"
	test "golang_project/test/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ContextMock struct {
	JSONCalled bool
}

//1.
func TestCreateFriendConnectionSuccessfulCase(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", strings.NewReader("{\"friends\":[\"fda@yahoo.com.vn\",\"hsa@s3corp.com.vn\"]}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.FriendListResponse{
		Success: true,
		Friends: nil,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestCreateFriendConnectionWithEmptyBody(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", strings.NewReader("{\"friends\":[]"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateFriendConnectionWithNoFriendEmail(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateFriendConnectionWithOnlyOneEmail(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", strings.NewReader("{\"friends\":[\"thehaohcm@yahoo.com.vn\"]"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//2.
func TestShowFriendsByEmailSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", strings.NewReader("{\"email\":\"thehaohcm@yahoo.com.vn\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.FriendListResponse{
		Success: true,
		Friends: []string{"hao.nguyen@s3corp.com.vn"},
		Count:   1,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowFriendsByEmailEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowFriendsByEmailWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

//3.
func TestShowCommonFriendListSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", strings.NewReader(`{
		"friends": [
		  "thehaohcm@yahoo.com.vn","chinh.nguyen@s3corp.com.vn"
		]
	  }`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.FriendListResponse{
		Success: true,
		Friends: []string{"hao.nguyen@s3corp.com.vn"},
		Count:   1,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowCommonFriendListEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowCommonFriendListWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

//4.
func TestSubscribeFromEmailSuccessfulCase(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", strings.NewReader(`{
	"requestor": "lisa@example.com",
	"target": "john@example.com"	
	}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.SubscribeResponse{
		Success: true,
	}
	var modelRes models.SubscribeResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestSubscribeFromEmailFailCaseEmptyTarget(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", strings.NewReader(`{
	"requestor": "lisa@example.com"	
	}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSubscribeFromEmailFailCaseEmptyRequestor(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", strings.NewReader(`{
	"target": "lisa@example.com"	
	}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSubscribeFromEmailFailCaseEmptyBody(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//5.
func TestBlockSubscribeByEmailSuccessfulCase(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", strings.NewReader(`{
		"requestor": "lisa@example.com",
		"target": "john@example.com"
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.BlockSubscribeResponse{
		Success: true,
	}
	var modelRes models.BlockSubscribeResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestBlockSubscribeByEmailFailCaseEmptyTarget(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", strings.NewReader(`{
		"requestor": "lisa@example.com"	
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBlockSubscribeByEmailFailCaseEmptyRequestor(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", strings.NewReader(`{
		"target": "lisa@example.com"	
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBlockSubscribeByEmailFailCaseEmptyBody(t *testing.T) {
	router := test.SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// 6.
func TestShowSubscribingEmailListByEmailSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", strings.NewReader("{\"sender\": \"thehaohcm@yahoo.com.vn\",\"text\": \"Hello World! kate@example.com\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	exRs := models.GetSubscribingEmailListResponse{
		Success:    true,
		Recipients: []string{"hao.nguyen@s3corp.com.vn", "kate@example.com"},
	}

	var modelRes models.GetSubscribingEmailListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowSubscribingEmailListByEmailEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowSubscribingEmailListByEmailWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}
