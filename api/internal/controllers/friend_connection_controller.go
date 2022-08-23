package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"
	"golang_project/api/internal/services"
)

// FriendConnectionController interface declares all functions used in Controller layer
type FriendConnectionController interface {
	CreateUser(c *gin.Context)
	CreateFriendConnection(c *gin.Context)
	GetFriendListByEmail(c *gin.Context)
	ShowCommonFriendList(c *gin.Context)
	SubscribeFromEmail(c *gin.Context)
	BlockSubscribeByEmail(c *gin.Context)
	GetSubscribingEmailListByEmail(c *gin.Context)
}

type controller struct {
	service services.FriendConnectionService
}

// New function used for initializing a FriendConnectionController
// pass a FriendConnectionService as parameter
func New(service services.FriendConnectionService) FriendConnectionController {
	return &controller{
		service: service,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Create an User
// @Schemes
// @Description Extend request: create a new user
// @Tags User API
// @Accept json
// @Produce json
// @Param   Request body models.CreatingUserRequest true "Create an User"
// @Router /users/createUser [post]
// CreateUser function works as a controller for creating an new user
// pass a gin's context as parameter
func (ctl *controller) CreateUser(c *gin.Context) {
	var request models.CreatingUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pkg.CheckValidEmail(request.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Create a friend connection
// @Schemes
// @Description Requirement 1: As a user, I need an API to create a friend connection between two email addresses.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.FriendConnectionRequest true "Create a friend connection between 2 user emails"
// @Router /friends/createConnection [post]
// CreateFriendConnection function works as a controller for creating friend connection between 2 user emails
// pass a gin's context as parameter
func (ctl *controller) CreateFriendConnection(c *gin.Context) {
	var request models.FriendConnectionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Friends) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request, the model must not be empty"})
		return
	}

	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.CreateConnection(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Get Friend list by email
// @Schemes
// @Description Requirement 2: As a user, I need an API to retrieve the friends list for an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.FriendListRequest true "Get a list of friend by user email"
// @Router /friends/showFriendsByEmail [post]
// GetFriendListByEmail function works as a controller for getting a friend list by an email address
// pass a gin's context as parameter
func (ctl *controller) GetFriendListByEmail(c *gin.Context) {
	var request models.FriendListRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request == (models.FriendListRequest{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request, the model must not be empty"})
		return
	}

	if err := pkg.CheckValidEmail(request.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.GetFriendConnection(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Show common Friend list
// @Schemes
// @Description Requirement 3: As a user, I need an API to retrieve the common friends list between two email addresses.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.CommonFriendListRequest true "Retrieve the common friends list between two email addresses"
// @Router /friends/showCommonFriendList [post]
// ShowCommonFriendList function works as a controller for getting a list of common friends between two email addresses
// pass a gin's context as parameter
func (ctl *controller) ShowCommonFriendList(c *gin.Context) {
	var request models.CommonFriendListRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Friends == nil || len(request.Friends) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request, the friends list must have over 1 item"})
		return
	}

	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.ShowCommonFriendList(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Create a subscribe from email
// @Schemes
// @Description Requirement 4: As a user, I need an API to subscribe to updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.SubscribeRequest true "Subscribe to updates from an email address"
// @Router /friends/subscribeFromEmail [post]
// SubscribeFromEmail function works as a controller for creating a subscribe from an email address to another one
// pass a gin's context as parameter
func (ctl *controller) SubscribeFromEmail(c *gin.Context) {
	var request models.SubscribeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Requestor == "" || request.Target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, both requestor and target must not be null"})
		return
	}

	if err := pkg.CheckValidEmails([]string{request.Requestor, request.Target}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.SubscribeFromEmail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Block subscribe by email
// @Schemes
// @Description Requirement 5: As a user, I need an API to block updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.BlockSubscribeRequest true "Block updates from an email address"
// @Router /friends/blockSubscribeByEmail [post]
// BlockSubscribeByEmail function works as a controller for creating a block subscribe update from an email address to another one
// pass a gin's context as parameter
func (ctl *controller) BlockSubscribeByEmail(c *gin.Context) {
	var request models.BlockSubscribeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Requestor == "" || request.Target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, both requestor and target must not be null"})
		return
	}

	if err := pkg.CheckValidEmails([]string{request.Requestor, request.Target}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.BlockSubscribeByEmail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary Get Subscribing email list by email
// @Schemes
// @Description Requirement 6: As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.GetSubscribingEmailListRequest true "retrieve all email addresses that can receive update from an email address"
// @Router /friends/showSubscribingEmailListByEmail [post]
// GetSubscribingEmailListByEmail function works as a controller for getting a list of subscribe email by an email address
// pass a gin's context as parameter
func (ctl *controller) GetSubscribingEmailListByEmail(c *gin.Context) {
	var request models.GetSubscribingEmailListRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Sender == "" || request.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, both Sender and Text must not be null"})
		return
	}

	if err := pkg.CheckValidEmail(request.Sender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ctl.service.GetSubscribingEmailListByEmail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
