package handler

import (
	"goapi/pkg/common"
	"goapi/pkg/module/auth/core/dto"
	"goapi/pkg/module/auth/core/ports"
)

type AuthHandler struct {
	serv ports.AuthService
}

func NewAuthHandler(serv ports.AuthService) *AuthHandler {
	return &AuthHandler{serv}
}

// @Summary Register a new user
// @Description Register a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body swagger.RegisterForm true "User Data"
// @Failure 422 {object} swagdto.Error422{error=swagger.ErrCreateSampleData}
// @Failure 500 {object} swagdto.Error500
// @Success 201
// @Router /auth/register [post]
func (h AuthHandler) Register(c common.HContext) error {
	// แปลง JSON เป็น struct
	form := new(dto.RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return common.ResponseError(c, common.ErrBodyParser)
	}
	// ส่งต่อไปให้ service ทำงาน
	err := h.serv.Register(*form, c.RequestId())
	if err != nil {
		// error จะถูกจัดการมาจาก service แล้ว
		return common.ResponseError(c, err)
	}

	return common.ResponseCreated(c, "", nil)
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body swagger.LoginForm true "Login Data"
// @Failure 422 {object} swagdto.Error422{error=swagger.ErrCreateSampleData}
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagdto.Response{data=swagger.AuthSampleData}
// @Router /auth [post]
func (h AuthHandler) Login(c common.HContext) error {
	// แปลง JSON เป็น struct
	form := new(dto.LoginForm)
	if err := c.BodyParser(form); err != nil {
		return common.ResponseError(c, common.ErrBodyParser)
	}
	// ส่งต่อไปให้ service ทำงาน
	auth, err := h.serv.Login(*form, c.RequestId())
	if err != nil {
		// error จะถูกจัดการมาจาก service แล้ว
		return common.ResponseError(c, err)
	}

	return common.ResponseOk(c, "auth", auth)
}

// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body swagger.LogoutForm true "Logout Data"
// @Failure 422 {object} swagdto.Error422{error=swagger.ErrCreateSampleData}
// @Failure 500 {object} swagdto.Error500
// @Success 200
// @Router /auth/logout [post]
func (h AuthHandler) Logout(c common.HContext) error {
	// แปลง JSON เป็น struct
	form := new(dto.LogoutForm)
	if err := c.BodyParser(form); err != nil {
		return common.ResponseError(c, common.ErrBodyParser)
	}
	// ส่งต่อไปให้ service ทำงาน
	err := h.serv.Logout(*form, c.RequestId())
	if err != nil {
		// error จะถูกจัดการมาจาก service แล้ว
		return common.ResponseError(c, err)
	}

	return common.ResponseOk(c, "", nil)
}

// @Summary Refresh Access Token
// @Description Refresh Access Token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body swagger.RefreshForm true "Logout Data"
// @Failure 422 {object} swagdto.Error422{error=swagger.ErrCreateSampleData}
// @Failure 500 {object} swagdto.Error500
// @Success 200
// @Router /auth/logout [post]
func (h AuthHandler) Refresh(c common.HContext) error {
	// แปลง JSON เป็น struct
	form := new(dto.RefreshForm)
	if err := c.BodyParser(form); err != nil {
		return common.ResponseError(c, common.ErrBodyParser)
	}
	// ส่งต่อไปให้ service ทำงาน
	auth, err := h.serv.Refresh(*form, c.RequestId())
	if err != nil {
		// error จะถูกจัดการมาจาก service แล้ว
		return common.ResponseError(c, err)
	}

	return common.ResponseOk(c, "auth", auth)
}

// @Summary Get a user profile
// @Description Get a specific user by id
// @Produce json
// @Tags Auth
// @Failure 401 {object} swagdto.Error401
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagdto.Response{data=swagger.UserSampleData}
// @Router /auth/profile [get]
func (h AuthHandler) Profile(c common.HContext) error {
	id := c.Param("id")

	user, err := h.serv.Profile(id, c.RequestId())

	if err != nil {
		return common.ResponseError(c, err)
	}

	return common.ResponseOk(c, "user", user)
}
