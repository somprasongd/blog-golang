package mapper

import (
	"goapi/pkg/module/users/core/dto"
	"goapi/pkg/module/users/core/model"
)

func CreateUserFormToModel(dto dto.NewUserForm) *model.User {
	return &model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func UserToDto(m *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    m.ID.String(),
		Email: m.Email,
		Role:  m.Role.String(),
	}
}

func UsersToDto(Users model.Users) dto.UserResponses {
	dtos := make([]*dto.UserResponse, len(Users))
	for i, t := range Users {
		dtos[i] = UserToDto(t)
	}

	return dtos
}
