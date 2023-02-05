package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr string
}

func Read() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, fmt.Errorf("load: %w", err)
	}

	cfg := Config{
		Addr: os.Getenv("GATE_PORT"),
	}
	if cfg.Addr == "" {
		return Config{}, errors.New("Addr is empty")
	}
	return cfg, nil
}
