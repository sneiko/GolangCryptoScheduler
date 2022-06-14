package app

import (
	"CryptoTest/internal/controllers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func loadControllers(app *fiber.App) {
	api := app.Group("/api")

	// Crypto
	api.Get("/cryptoHistory/:pair", controllers.CryptoLast)
	api.Post("/cryptoHistory/:pair", controllers.CryptoHistory)

	// Fiat
	api.Get("/fiatHistory/:pair", controllers.FiatLast)
	api.Post("/fiatHistory/:pair", controllers.FiatHistory)
}

func setupSwagger(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
