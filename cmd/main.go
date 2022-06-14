package main

import (
	"CryptoTest/internal/app"
	"github.com/caarlos0/env"
	"log"
)

// @title CryptoTest API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name Neykovich Sergey
// @contact.email s_neiko@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @BasePath /
func main() {
	cfg := app.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	app.Run(cfg)
}
