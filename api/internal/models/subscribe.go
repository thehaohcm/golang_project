package models

// SubscribeRequest struct used when user request the service to create a subscribe
type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// SubscribeResponse struct used when the service response a status after creating a subscribe
type SubscribeResponse struct {
	Success bool `json:"success"`
}

// BlockSubscribeRequest struct used when user request the service to block an existing subscribe
type BlockSubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// BlockSubscribeResponse struct used when the service response a status after blocking a subscribe
type BlockSubscribeResponse struct {
	Success bool `json:"success"`
}

// GetSubscribingEmailListRequest struct used when user request the service to get list of subscribe emails
type GetSubscribingEmailListRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// GetSubscribingEmailListResponse struct used when the service response a list of emails
type GetSubscribingEmailListResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
