package dto

type NewUserForm struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserId struct {
	ID string `validate:"required,uuid4"`
}

type UpdateUserPasswordForm struct {
	PasswordOld string `json:"password_old"`
	PasswordNew string `json:"password_new"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserResponses []*UserResponse
