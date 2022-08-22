package models

// CommonFriendListRequest struct used when user request the service to get a common friend list
// Friends []string property stores a list of user emails
type CommonFriendListRequest struct {
	Friends []string `json:"friends"`
}

// CommonFriendListResponse struct used when the service return a common friend list
// Success bool property store the status of API process
// Friends []string property stores a list of common friend emails
// Count int property store a number of common friend emails
type CommonFriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}
