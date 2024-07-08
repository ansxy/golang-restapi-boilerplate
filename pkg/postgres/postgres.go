package custome_postgres

import (
	"github.com/ansxy/golang-boilerplate-gin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient(conf *config.Config) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(conf.Postgres.URI), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
