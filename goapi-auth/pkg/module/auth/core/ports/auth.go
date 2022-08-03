package ports

import (
	"goapi/pkg/module/auth/core/dto"
	"goapi/pkg/module/user/core/model"
)

// interface สำหรับ output port
type AuthRepository interface {
	FindUserByEmail(email string) (*model.User, error)
	CreateUser(*model.User) error
}

// interface สำหรับ input port
type AuthService interface {
	Register(form dto.RegisterForm, reqId string) error
	Login(form dto.LoginForm, reqId string) (*dto.AuthResponse, error)
	Profile(email string, reqId string) (*dto.UserInfo, error)
}
