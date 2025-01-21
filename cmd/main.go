package main

import (
	"fmt"
	"log"

	"github.com/erknas/wt-weapons/internal/config"
	"github.com/erknas/wt-weapons/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
