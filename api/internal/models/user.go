package models

// CreatingUserRequest struct used when the service return a process status after creating a user
type CreatingUserRequest struct {
	Email string `json:"email"`
}

// CreatingUserResponse struct used when the service return a process status after creating a user
type CreatingUserResponse struct {
	Success bool `json:"success"`
}

// User struct used when mapping to get a User model after querying data from User table in database
type User struct {
	Email string
}
