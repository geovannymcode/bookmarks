package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetDb(config AppConfig) *pgx.Conn {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.UserName, config.Db.Password, config.Db.Database)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
