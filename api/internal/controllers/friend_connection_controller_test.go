package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"golang_project/api/internal/docs"
	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"
)

type ContextMock struct {
	JSONCalled bool
}

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
	assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"Invalid Request, the model must not be empty\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"Invalid email address\"}", w.Body.String())
}

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
	assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"Invalid Request, the friends list must have over 1 item\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"invalid character 'c' after array element\"}", w.Body.String())
}

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
	assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"Invalid request, both Sender and Text must not be null\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"Invalid email address\"}", w.Body.String())
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) CreateUser(req models.CreatingUserRequest) (models.CreatingUserResponse, error) {
	if err := pkg.CheckValidEmail(req.Email); err != nil {
		return models.CreatingUserResponse{}, err
	}
	return models.CreatingUserResponse{Success: true}, nil
}

func (s *ServiceMock) CreateConnection(request models.FriendConnectionRequest) (models.FriendConnectionResponse, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return models.FriendConnectionResponse{Success: true}, errors.New("Invalid email address")
	}
	if len(request.Friends) > 0 {
		return models.FriendConnectionResponse{Success: true}, nil
	}
	return models.FriendConnectionResponse{Success: false}, nil
}
func (s *ServiceMock) GetFriendConnection(request models.FriendListRequest) (models.FriendListResponse, error) {
	if err := pkg.CheckValidEmail(request.Email); err != nil {
		return models.FriendListResponse{Success: false}, err
	}
	return models.FriendListResponse{Success: true, Friends: []string{"thehaohcm@yahoo.com.vn", "hao.nguyen@s3corp.com.vn"}, Count: 2}, nil
}
func (s *ServiceMock) ShowCommonFriendList(request models.CommonFriendListRequest) (models.CommonFriendListResponse, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return models.CommonFriendListResponse{}, err
	}
	if len(request.Friends) > 0 {
		return models.CommonFriendListResponse{Success: true, Friends: []string{"hao.nguyen@s3corp.com.vn"}, Count: 1}, nil
	}
	return models.CommonFriendListResponse{}, nil
}
func (s *ServiceMock) SubscribeFromEmail(request models.SubscribeRequest) (models.SubscribeResponse, error) {
	if err := pkg.CheckValidEmails([]string{request.Requestor, request.Target}); err != nil {
		return models.SubscribeResponse{}, err
	}
	if request.Requestor != "" && request.Target != "" {
		return models.SubscribeResponse{Success: true}, nil
	}
	return models.SubscribeResponse{}, nil
}
func (s *ServiceMock) BlockSubscribeByEmail(request models.BlockSubscribeRequest) (models.BlockSubscribeResponse, error) {
	if err := pkg.CheckValidEmail(request.Requestor); err != nil {
		return models.BlockSubscribeResponse{}, err
	}
	if request.Requestor != "" && request.Target != "" {
		return models.BlockSubscribeResponse{Success: true}, nil
	}
	return models.BlockSubscribeResponse{}, nil
}
func (s *ServiceMock) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) (models.GetSubscribingEmailListResponse, error) {
	if err := pkg.CheckValidEmail(request.Sender); err != nil {
		return models.GetSubscribingEmailListResponse{}, err
	}
	if request.Text != "" {
		if request.Sender == "thehaohcm@yahoo.com.vn" && request.Text == "Hello World! kate@example.com" {
			return models.GetSubscribingEmailListResponse{Success: true, Recipients: []string{"hao.nguyen@s3corp.com.vn", "kate@example.com"}}, nil
		}
		return models.GetSubscribingEmailListResponse{Success: true, Recipients: []string{"abc@gmail.com"}}, nil
	}
	return models.GetSubscribingEmailListResponse{}, nil
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
			v1.POST("/users/createUser", controller.CreateUser)

			v1.POST("/friends/createConnection", controller.CreateFriendConnection)

			v1.POST("/friends/showFriendsByEmail", controller.GetFriendListByEmail)

			v1.POST("/friends/showCommonFriendList", controller.ShowCommonFriendList)

			v1.POST("/friends/subscribeFromEmail", controller.SubscribeFromEmail)

			v1.POST("/friends/blockSubscribeByEmail", controller.BlockSubscribeByEmail)

			v1.POST("/friends/showSubscribingEmailListByEmail", controller.GetSubscribingEmailListByEmail)
		}
	}
	return router
}
