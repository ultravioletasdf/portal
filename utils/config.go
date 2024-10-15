package utils

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Address string `env:"ADDRESS,notEmpty" envDefault:":3000"`
	Prefork bool   `env:"PREFORK" envDefault:"true"`
  Development bool `env:"DEVELOPMENT" envDefault:"False"`

  TursoAddress string `env:"TURSO_ADDRESS,required,notEmpty"`
  TursoToken string `env:"TURSO_TOKEN,required,notEmpty"`
  DatabaseFileName string `env:"DATABASE_FILE_NAME,notEmpty" envDefault:"app.db"`
}

func LoadConfig() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
	return &cfg
}
