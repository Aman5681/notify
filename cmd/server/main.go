package main

import (
	"log"

	"github.com/Aman5681/notify/internal/api"
	"github.com/Aman5681/notify/internal/config"
	"github.com/Aman5681/notify/internal/db"
)

func main() {
	config := config.LoadConfig()
	db.InitDB(&config)
	r := api.SetupRouter()
	log.Println("Starting DevNotify server on :8080")
	r.Run(":8080")
}
