package models

// FriendListRequest struct used when user request the service to get a list of friend emails
// Email string property stores a list of friend emails
type FriendListRequest struct {
	Email string `json:"email"`
}

// FriendListResponse struct used when the service return a list of friend emails
// Success bool property stores a status of API process
// Friends []string property stores a list of friend emails
// Count int property store a number of common friend emails
type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
