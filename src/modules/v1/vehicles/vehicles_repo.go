package vehicles

import (
	"errors"

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
	result := r.db.Where("vehicles_id", params).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Menghapus data")
	}

	return data, nil
}
func (r *vehicles_repo) Update(data *models.Vehicle, params string) (*models.Vehicle, error) {
	result := r.db.Model(&data).Where("vehicles_id = ?", params).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Mengupdate data")
	}
	// id := strconv.FormatUint(uint64(data.Vehicles_id), 10)
	// if id == params {
	// 	return nil, errors.New("id sudah terpakai")
	// }
	r.db.Create(&data)
	return data, nil
}
