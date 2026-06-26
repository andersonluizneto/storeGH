package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/andersonluizneto/storeGH/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}
