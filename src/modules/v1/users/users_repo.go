package users

import (
	"errors"
	"os"

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

func (r *users_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "username = ?", username)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *users_repo) FindByRole(role string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "role = ?", role)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}
func (r *users_repo) UserExsist(username, email string) bool {
	var data models.User
	result := r.db.First(&data, "username = ? OR email = ?", username, email)
	if result.Error != nil {
		return false
	}
	return true
}

func (r *users_repo) Save(data *models.User) (*models.User, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal Menyimpan data")
	}

	return data, nil
}
func (r *users_repo) Delete(data *models.User, params string) (*models.User, error) {

	//delete file
	r.db.First(&data, "user_id = ?", params)
	path := "./image/" + data.Image
	err := os.Remove(path)
	if err != nil {
		return nil, errors.New("file tidak ada")
	}
	//delete data
	result := r.db.Where("user_id", params).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Menghapus data")
	}

	return data, nil
}

func (r *users_repo) Update(data *models.User, params string, filename string) (*models.User, error) {

	r.db.First(&data, "user_id = ?", params)
	path := "./image/" + data.Image

	if data.Image != "" {
		os.Remove(path)
	}
	data.Image = filename
	result := r.db.Model(&data).Where("user_id = ?", params).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Mengupdate data")
	}

	return data, nil
}
