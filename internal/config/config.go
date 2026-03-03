package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ReinaConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`

	DebugMode bool `yaml:"debugMode"`
}

func LoadConfig() *ReinaConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file! %q", err)
	}

	address := os.Getenv("REINA_ADDRESS")
	port, err := strconv.Atoi(os.Getenv("REINA_PORT"))
	if err != nil {
		log.Fatalf("Failed to format port from .env!")
	}

	isDebug, err := strconv.ParseBool(os.Getenv("REINA_IS_DEBUG"))
	if err != nil {
		log.Fatalf("Failed to format IsDebug from .env!")
	}

	return &ReinaConfig{
		Address:   address,
		Port:      port,
		DebugMode: isDebug,
	}
}
