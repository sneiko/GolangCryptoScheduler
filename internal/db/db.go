package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
	"path/filepath"
)

var Db *pgx.Conn

func Connect(dbUrl string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}

func LoadScheme() error {
	path := filepath.Join("./internal/db/scheme/init_tables.sql")

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)
	_, err := Db.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	return nil
}
