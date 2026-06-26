package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andersonluizneto/storeGH/configs"
	"github.com/andersonluizneto/storeGH/internal/shared/database"
)

func main() {
	configs.Load()

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server_port := fmt.Sprintf(":%s", configs.GetServerPort())
	err = http.ListenAndServe(server_port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
