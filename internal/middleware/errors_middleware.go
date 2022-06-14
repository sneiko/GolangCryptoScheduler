package middleware

import (
	"CryptoTest/internal/db/entity"
	"github.com/gofiber/fiber/v2"
)

var ErrorMiddleware = func(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	responseObj := entity.NewError(err.Error())

	return ctx.Status(code).JSON(responseObj)
}
