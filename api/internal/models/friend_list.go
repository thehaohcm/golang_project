package models

// FriendListRequest struct used when user request the service to get a list of friend emails
type FriendListRequest struct {
	Email string `json:"email"`
}

// FriendListResponse struct used when the service return a list of friend emails
type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
