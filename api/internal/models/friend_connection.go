package models

// FriendConnectionRequest struct used when user request the service to get a list of friend emails
type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

// FriendConnectionResponse struct used when the service return a list of friend emails
type FriendConnectionResponse struct {
	Success bool `json:"success"`
}
