package main

import (
	"fmt"
	"log"

	"github.com/ElladanTasartir/fiber-auth/internal/config"
	"github.com/ElladanTasartir/fiber-auth/internal/storage"
)

func main() {
	config, err := config.NewConfig("./config")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.DB.Exec("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
