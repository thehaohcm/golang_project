package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang_project/api/internal/docs"
	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ContextMock struct {
	JSONCalled bool
}

// external
func TestCreateUserSuccessfulCase(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/createUser", strings.NewReader("{\"email\":\"fda@yahoo.com.vn\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.CreatingUserResponse{Success: true}
	var modelRes models.CreatingUserResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestCreateUserFailCaseWithInvalidEmail(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/createUser", strings.NewReader("{\"email\":\"fda\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserFailCaseWithEmptyEmail(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/createUser", strings.NewReader("{\"email\":\"\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserFailCaseWithNilBody(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/createUser", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// 1.
func TestCreateFriendConnectionSuccessfulCase(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", strings.NewReader("{\"friends\":[\"thehaohcm@yahoo.com.vn\"]"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateFriendConnectionWithInvalidEmail(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/createConnection", strings.NewReader("{\"friends\":[\"thehaohcm\"]"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// 2.
func TestShowFriendsByEmailSuccessfulCode(t *testing.T) {
	router := SetupRouterForTesting()

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
		Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"},
		Count:   2,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowFriendsByEmailEmptyBody(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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

func TestShowFriendsByEmailWithInvalidEmail(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", strings.NewReader("{\"email\":\"thehaohcm\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

// 3.
func TestShowCommonFriendListSuccessfulCode(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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

func TestShowCommonFriendListWithInvalid(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", strings.NewReader(`{
		"friends": [
		  "thehaohcm,"chinh.nguyen@s3corp.com.vn"
		]
	  }`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

// 4.
func TestSubscribeFromEmailSuccessfulCase(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSubscribeFromEmailWithInvalidEmailTarget(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", strings.NewReader(`{
		"requestor": "lisa@example.com",
		"target": "john"	
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSubscribeFromEmailWithInvalidEmailRequestor(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/subscribeFromEmail", strings.NewReader(`{
		"requestor": "lisa",
		"target": "john@example.com"	
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// 5.
func TestBlockSubscribeByEmailSuccessfulCase(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBlockSubscribeByEmailFailCaseWithInvalidEmailRequestor(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", strings.NewReader(`{
		"requestor": "lisa",
		"target": "john@example.com"
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBlockSubscribeByEmailFailCaseWithInvalidEmailTarget(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/blockSubscribeByEmail", strings.NewReader(`{
		"requestor": "lisa@example.com",
		"target": "john"
		}`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// 6.
func TestShowSubscribingEmailListByEmailSuccessfulCode(t *testing.T) {
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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
	router := SetupRouterForTesting()

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

func TestShowSubscribingEmailListByEmailWithInvalidEmail(t *testing.T) {
	router := SetupRouterForTesting()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", strings.NewReader("{\"sender\": \"thehaohcm\",\"text\": \"Hello World! kate@example.com\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) CreateUser(req models.CreatingUserRequest) models.CreatingUserResponse {
	if valid, err := pkg.CheckValidEmail(req.Email); !valid || err != nil {
		return models.CreatingUserResponse{}
	}
	return models.CreatingUserResponse{Success: true}
}

func (s *ServiceMock) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	if len(request.Friends) > 0 {
		return models.FriendConnectionResponse{Success: true}
	}
	return models.FriendConnectionResponse{Success: false}
}
func (s *ServiceMock) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	if request.Email == "" {
		return models.FriendListResponse{Success: false}
	}
	return models.FriendListResponse{Success: true, Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}, Count: 2}
}
func (s *ServiceMock) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	if valid, err := pkg.CheckValidEmails(request.Friends); !valid || err != nil {
		panic("invalid email address")
	}
	if len(request.Friends) > 0 {
		return models.CommonFriendListResponse{Success: true, Friends: []string{"hao.nguyen@s3corp.com.vn"}, Count: 1}
	}
	return models.CommonFriendListResponse{}
}
func (s *ServiceMock) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	if request.Requestor != "" && request.Target != "" {
		return models.SubscribeResponse{Success: true}
	}
	return models.SubscribeResponse{}
}
func (s *ServiceMock) BlockSubscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	if request.Requestor != "" && request.Target != "" {
		return models.BlockSubscribeResponse{Success: true}
	}
	return models.BlockSubscribeResponse{}
}
func (s *ServiceMock) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	if request.Sender != "" && request.Text != "" {
		if request.Sender == "thehaohcm@yahoo.com.vn" && request.Text == "Hello World! kate@example.com" {
			return models.GetSubscribingEmailListResponse{Success: true, Recipients: []string{"hao.nguyen@s3corp.com.vn", "kate@example.com"}}
		}
		return models.GetSubscribingEmailListResponse{Success: true, Recipients: []string{"abc@gmail.com"}}
	}
	return models.GetSubscribingEmailListResponse{}
}

func SetupRouterForTesting() *gin.Engine {
	serv := &ServiceMock{}
	controller := New(serv)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//external
			v1.POST("/users/createUser", controller.CreateUser)

			//1. Done
			v1.POST("/friends/createConnection", controller.CreateFriendConnection)

			//2. Done
			v1.POST("/friends/showFriendsByEmail", controller.GetFriendListByEmail)

			//3. Done
			v1.POST("/friends/showCommonFriendList", controller.ShowCommonFriendList)

			//4. Done
			v1.POST("/friends/subscribeFromEmail", controller.SubscribeFromEmail)

			//5. Done
			v1.POST("/friends/blockSubscribeByEmail", controller.BlockSubscribeByEmail)

			//6. Done
			v1.POST("/friends/showSubscribingEmailListByEmail", controller.GetSubscribingEmailListByEmail)
		}
	}
	return router
}
