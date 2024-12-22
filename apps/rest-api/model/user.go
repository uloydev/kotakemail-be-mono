package model

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
