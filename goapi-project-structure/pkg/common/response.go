package common

import (
	"goapi-project-structure/pkg/common/errs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status    int                    `json:"status" example:"200"`
	Data      map[string]interface{} `json:"data,omitempty" example:"{data:{task}}"`
	Error     interface{}            `json:"error,omitempty" example:"{}"`
	RequestId string                 `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}

func ResponseCreated(c *fiber.Ctx, key string, body interface{}) error {
	res := Response{
		Status:    http.StatusCreated,
		Data:      map[string]interface{}{key: body},
		Error:     nil,
		RequestId: c.GetRespHeader("X-Request-ID"),
	}

	c.Status(http.StatusCreated)
	return c.JSON(res)
}

func ResponseError(c *fiber.Ctx, err error) error {
	var appErr errs.AppError
	switch e := err.(type) {
	case errs.AppError:
		appErr = e
	default: // case error
		appErr = errs.AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	res := Response{
		Status:    appErr.Code,
		Data:      nil,
		Error:     appErr,
		RequestId: c.GetRespHeader("X-Request-ID"),
	}
	c.Status(appErr.Code)
	return c.JSON(res)
}
