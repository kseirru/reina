package main

import (
	"ReinaMusic/internal/app"
	"ReinaMusic/internal/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	if cfg == nil {
		log.Fatalf("Failed to load config!")
		return
	}

	if err := app.InitReinaRest(cfg); err != nil {
		log.Fatalf("Failed to run Reina REST server!")
	}
}
