package dto

// Response untuk detail user
type UserDetailResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response untuk list user
type UserListResponse struct {
	Users []UserDetailResponse `json:"users"`
}

// Response error standar
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
