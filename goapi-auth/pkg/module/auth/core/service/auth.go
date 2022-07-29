package service

import (
	"errors"
	"goapi/pkg/common"
	"goapi/pkg/common/logger"
	"goapi/pkg/module/auth/core/dto"
	"goapi/pkg/module/auth/core/ports"
	"goapi/pkg/module/users/core/model"
	userPorts "goapi/pkg/module/users/core/ports"
	"goapi/pkg/util"
)

var (
	// ErrUserNotFoundById auth not found error when find with id
	ErrUserNotFoundById     = common.NewNotFoundError("auth with given id not found")
	ErrHashPassword         = common.NewUnexpectedError("hash password error")
	ErrUserEmailDuplication = common.NewBadRequestError("duplicate email")
	ErrUserPasswordNotMatch = common.NewBadRequestError("password is not macth")
)

type authService struct {
	repo userPorts.UserRepository
}

func NewUserService(repo userPorts.UserRepository) ports.AuthService {
	return &authService{repo}
}

func (s authService) Register(form dto.RegisterForm, reqId string) error {
	// validate
	if err := common.ValidateDto(form); err != nil {
		return common.NewInvalidError(err.Error())
	}

	u, err := s.repo.FindByEmail(form.Email)

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		logger.ErrorWithReqId(err.Error(), reqId)
		return common.ErrDbQuery
	}

	if u != nil {
		return ErrUserEmailDuplication
	}

	auth := model.User{Email: form.Email}
	hashPwd, err := util.HashPassword(form.Password)
	if err != nil {
		logger.ErrorWithReqId(err.Error(), reqId)
		return ErrHashPassword
	}
	auth.Password = hashPwd

	err = s.repo.Create(&auth)
	if err != nil {
		logger.ErrorWithReqId(err.Error(), reqId)
		return common.ErrDbInsert
	}

	return nil
}

func (s authService) Login(form dto.LoginForm, reqId string) (*dto.AuthResponse, error) {
	// validate
	err := common.ValidateDto(form)
	if err != nil {
		return nil, common.NewInvalidError(err.Error())
	}

	user, err := s.repo.FindByEmail(form.Email)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, ErrUserNotFoundById
		}
		logger.ErrorWithReqId(err.Error(), reqId)
		return nil, common.ErrDbQuery
	}

	match := util.CheckPasswordHash(form.Password, user.Password)

	if !match {
		return nil, ErrUserPasswordNotMatch
	}

	// TODO: Gen Refresh Token
	// TODO: Gen Access Token

	serialized := dto.AuthResponse{
		User: dto.UserInfo{
			ID:    user.ID.String(),
			Email: user.Email,
			Role:  user.Role.String(),
		},
	}
	return &serialized, nil
}

func (s authService) Logout(form dto.LogoutForm, reqId string) error {
	return nil
}

func (s authService) Refresh(form dto.RefreshForm, reqId string) (*dto.AuthResponse, error) {
	return nil, nil
}

func (s authService) Profile(userId string, reqId string) (*dto.UserInfo, error) {
	return nil, nil
}
