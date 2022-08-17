package models

type CreatingUserRequest struct {
	Email string `json:"email"`
}

type CreatingUserResponse struct {
	Success bool `json:"success"`
}

type User struct {
	Email string
}
