package vehicles

import (
	"errors"
	"os"
	"strings"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"gorm.io/gorm"
)

// repo untuk komunikasi ke database

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db: db}
}

func (r *vehicles_repo) FindAll() (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *vehicles_repo) Save(data *models.Vehicle) (*models.Vehicle, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal Menyimpan data")
	}

	return data, nil
}

func (r *vehicles_repo) Delete(data *models.Vehicle, params string) (*models.Vehicle, error) {
	//delete file
	r.db.First(&data, "vehicles_id = ?", params)
	path := "./image/" + data.Image
	err := os.Remove(path)
	if err != nil {
		return nil, errors.New("file tidak ada")
	}

	result := r.db.Where("vehicles_id", params).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Menghapus data")
	}

	return data, nil
}
func (r *vehicles_repo) Update(data *models.Vehicle, params string, filename string) (*models.Vehicle, error) {
	r.db.First(&data, "vehicles_id = ?", params)
	path := "./image/" + data.Image

	if data.Image != "" {
		os.Remove(path)
	}
	data.Image = filename
	result := r.db.Model(&data).Where("vehicles_id = ?", params).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Mengupdate data")
	}
	return data, nil
}

//sort and search
func (r *vehicles_repo) Sort(params string) (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Order(params).Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}
func (r *vehicles_repo) Search(params string) (*models.Vehicles, error) {
	var data models.Vehicles
	vars := strings.ToLower((params))
	result := r.db.Where("LOWER(name) LIKE ?", "%"+vars+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mencari data")
	}

	return &data, nil
}

//popular vehicles
func (r *vehicles_repo) Popular() (*models.Vehicles, error) {
	var data models.Vehicles

	result := r.db.Order("orders desc").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}
