package models

// CommonFriendListRequest struct used when user request the service to get a common friend list
type CommonFriendListRequest struct {
	Friends []string `json:"friends"`
}

// CommonFriendListResponse struct used when the service return a common friend list
type CommonFriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
