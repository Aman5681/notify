package main

import (
	"log"

	"github.com/Aman5681/notify/internal/api"
)

func main() {
	r := api.SetupRouter()
	log.Println("Starting DevNotify server on :8080")
	r.Run(":8080")
}
