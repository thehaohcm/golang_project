package models

// SubscribeRequest struct used when user request the service to create a subscribe
// Requestor string property stores a requester email
// Target string property stores a target email
type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// SubscribeResponse struct used when the service response a status after creating a subscribe
// Success bool property stores a status of API process
type SubscribeResponse struct {
	Success bool `json:"success"`
}

// BlockSubscribeRequest struct used when user request the service to block an existing subscribe
// Requestor string property stores a requester email
// Target string property stores a target email
type BlockSubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// BlockSubscribeResponse struct used when the service response a status after blocking a subscribe
// Success bool property stores a status of API process
type BlockSubscribeResponse struct {
	Success bool `json:"success"`
}

// GetSubscribingEmailListRequest struct used when user request the service to get list of subscribe emails
// Sender string property stores a sender email
// Text string property stores a text
type GetSubscribingEmailListRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// GetSubscribingEmailListResponse struct used when the service response a list of emails
// Success bool property stores a status of API process
// Recipients []string property stores a list of emails
type GetSubscribingEmailListResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
