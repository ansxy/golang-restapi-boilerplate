package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
	BaseRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
