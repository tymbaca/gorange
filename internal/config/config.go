// Package config contains app configuration type.
package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Conf represents application config.
type Conf struct {
	PostgresHost     string `default:"localhost" envconfig:"POSTGRES_HOST"     required:"true"`
	PostgresPort     string `default:"5432"      envconfig:"POSTGRES_PORT"     required:"true"`
	PostgresDB       string `default:"postgres"  envconfig:"POSTGRES_DB"       required:"true"`
	PostgresUser     string `default:"postgres"  envconfig:"POSTGRES_USER"     required:"true"`
	PostgresPassword string `default:"postgres"  envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresMaxConn  int    `default:"10"        envconfig:"POSTGRES_MAX_CONN"`
}

func New() Conf {
	var cfg Conf
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
