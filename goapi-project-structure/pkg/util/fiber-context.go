package util

import (
	"goapi-project-structure/pkg/common"

	"github.com/gofiber/fiber/v2"
)

type fiberContext struct {
	*fiber.Ctx
}

func newFiberContext(c *fiber.Ctx) common.HContext {
	return &fiberContext{
		Ctx: c,
	}
}

func (c *fiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *fiberContext) BindQuery(v interface{}) error {
	return c.Ctx.QueryParser(v)
}

func (c *fiberContext) Query(key string) (string, bool) {
	q := c.Ctx.Query(key)
	return q, true
}

func (c *fiberContext) DefaultQuery(key string, d string) string {
	return c.Ctx.Query(key, d)
}

func (c *fiberContext) Param(key string) string {
	return c.Ctx.Params(key)
}

func (c *fiberContext) Header(key string) string {
	return c.Ctx.GetRespHeader(key)
}

func (c *fiberContext) Authorization() string {
	return c.Header("Authorization")
}

func (c *fiberContext) RequestId() string {
	return c.Header("x-trace-id")
}

func (c *fiberContext) ResponseJSON(code int, data interface{}) {
	c.Ctx.SendStatus(code)
	c.Ctx.JSON(data)
}

func WrapFiberHandler(h common.HandleFunc) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		h(newFiberContext(c))
		return nil
	}
}
