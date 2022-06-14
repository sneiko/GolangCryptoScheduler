package app

import (
	_ "CryptoTest/cmd/docs"
	"CryptoTest/internal/db"
	"CryptoTest/internal/middleware"
	"context"
	"fmt"
	"github.com/gkampitakis/fiber-modules/gracefulshutdown"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v4"
	"github.com/jasonlvhit/gocron"
	"log"
)

func Run(cfg Config) {
	defer func(Db *pgx.Conn, ctx context.Context) {
		if !Db.IsClosed() {
			err := Db.Close(ctx)
			if err != nil {
				log.Fatalf("%f", err)
			}
		}
	}(db.Db, context.Background())
	db.Db = db.Connect(cfg.DbConnectionString)

	err := db.LoadScheme()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorMiddleware})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Crypto API")
	})

	loadRepositories()
	preloadDataCheck()
	loadControllers(app)
	setupSwagger(app)
	SetupSchedulerJobs()

	apiUrl := fmt.Sprintf(":%d", cfg.Port)
	gracefulshutdown.Listen(app, apiUrl, gracefulshutdown.WithShutdownFns([]func() error{
		func() error {
			gocron.Clear()
			return db.Db.Close(context.Background())
		},
	}))
}
