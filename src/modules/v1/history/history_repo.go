package history

import (
	"errors"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db: db}
}

func (r *history_repo) FindAll() (*models.HistoriesAll, error) {
	var data models.HistoriesAll

	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *history_repo) Save(data *models.Histories) (*models.Histories, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal Menyimpan data")
	}

	return data, nil
}

func (r *history_repo) Delete(data *models.Histories, params string) (*models.Histories, error) {
	result := r.db.Where("history_id", params).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Menghapus data")
	}

	return data, nil
}
func (r *history_repo) Update(data *models.Histories, params string) (*models.Histories, error) {
	result := r.db.Model(&data).Where("history_id = ?", params).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("gagal Mengupdate data")
	}
	r.db.Create(&data)
	return data, nil
}

//sort and search
func (r *history_repo) Sort(params string) (*models.HistoriesAll, error) {
	var data models.HistoriesAll

	result := r.db.Order(params).Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}
func (r *history_repo) Search(params string) (*models.Histories, error) {
	var data models.Histories

	result := r.db.Where("history_id", params).Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mencari data")
	}

	return &data, nil
}
