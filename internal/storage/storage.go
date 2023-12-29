package storage

import (
	"fmt"

	"github.com/ElladanTasartir/fiber-auth/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(config *config.Config) (*Storage, error) {
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?%s",
		config.DB.User,
		config.DB.Password,
		config.DB.URL,
		config.DB.Database,
		config.DB.Options,
	)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db. err = %v", err)
	}

	return &Storage{
		DB: db,
	}, nil
}
