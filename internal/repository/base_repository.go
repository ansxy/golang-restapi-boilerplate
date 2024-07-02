package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct{}

func (repo *BaseRepository) Create(db *gorm.DB, data interface{}) error {
	err := db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) Update(db *gorm.DB, data interface{}) error {
	err := db.Model(data).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) Delete(db *gorm.DB, data interface{}) error {
	err := db.Delete(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *BaseRepository) FindOne(db *gorm.DB, data interface{}) (interface{}, error) {
	err := db.First(data).Preload(clause.Associations).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *BaseRepository) FindAll(db *gorm.DB, data interface{}) (interface{}, error) {
	err := db.Find(data).Preload(clause.Associations).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}
