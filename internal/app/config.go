package app

type Config struct {
	DbConnectionString string `env:"DB_CONNECTION_STRING"`
	Port               int    `env:"PORT" envDefault:"8080"`
	IsProduction       bool   `env:"IS_PRODUCTION" envDefault:"false"`
}
