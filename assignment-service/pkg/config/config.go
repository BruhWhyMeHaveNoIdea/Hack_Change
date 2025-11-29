package config

import (
	"hack_change/internal/auth"
	postgres "hack_change/pkg/db"

	"log"
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Server struct {
	RestPort int `env:"GRPC_REST_SERVER_PORT" env-default:"8080"`
}

type Config struct {
	Postgres postgres.Config
	JWT      auth.Config
}

func LoadConfig() (*Config, error) {
	// allow explicit path via DOTENV_PATH env var
	if p := os.Getenv("DOTENV_PATH"); p != "" {
		if err := godotenv.Load(p); err != nil {
			log.Printf("No .env found at DOTENV_PATH=%s: %v, falling back to search\n", p, err)
		} else {
			log.Printf("Loaded .env from DOTENV_PATH=%s\n", p)
		}
	} else {
		// try several likely locations (current dir, parent dirs)
		tried := []string{".env", "../.env", "../../.env", "../../../.env"}
		loaded := false
		for _, t := range tried {
			if _, err := os.Stat(t); err == nil {
				if err := godotenv.Load(t); err == nil {
					log.Printf("Loaded .env from %s\n", t)
					loaded = true
					break
				}
			}
		}
		if !loaded {
			log.Println("No .env file found in standard locations, using environment variables")
		}
	}

	cfg := &Config{}

	if err := env.Parse(&cfg.Postgres); err != nil {
		return nil, err
	}

	if err := env.Parse(&cfg.JWT); err != nil {
		return nil, err
	}

	return cfg, nil
}
