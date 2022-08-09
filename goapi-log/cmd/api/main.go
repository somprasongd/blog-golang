package main

import (
	"fmt"
	"goapi/pkg/common/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		log := c.Locals("log").(logger.Logger)
		log.Info("log in handler")
		return c.SendString("Hello, World!")
	})

	logger.Default.Info("start on port 3000")

	app.Listen(":3000")
}

func logMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fileds := map[string]interface{}{}
	fileds["requestid"] = c.GetRespHeaders()["X-Request-Id"]

	log := logger.NewWithFields(fileds)

	c.Locals("log", log)

	c.Next()

	// "[ip]:port status - method path (duration)"
	log.Info(fmt.Sprintf("[%v]:%v %v - %v %v (%v)", c.IP(), c.Port(), c.Response().StatusCode(), c.Method(), c.Path(), time.Since(start)))
	return nil
}
