package users

import (
	"errors"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"gorm.io/gorm"
)

//query ke database

type users_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *users_repo {
	return &users_repo{db: db}
}

func (r *users_repo) FindAll() (*models.Users, error) {
	var data models.Users

	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *users_repo) Save(data *models.User) (*models.User, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return data, nil
}
