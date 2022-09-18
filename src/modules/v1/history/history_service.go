package history

import (
	"fmt"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
)

type history_service struct {
	repo interfaces.HistoryRepo
}

func NewService(reps interfaces.HistoryRepo) *history_service {
	return &history_service{reps}
}

func (r *history_service) GetAll() (*models.HistoriesAll, error) {
	data, err := r.repo.FindAll()
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *history_service) Add(data *models.Histories) (*models.Histories, error) {
	data, err := r.repo.Save(data)
	if data != nil {
		fmt.Println("Data Berhasil Disimpan")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *history_service) DeleteData(data *models.Histories, params string) (*models.Histories, error) {
	data, err := r.repo.Delete(data, params)
	if data != nil {
		fmt.Println("Data Terhapus")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *history_service) UpdateData(data *models.Histories, params string) (*models.Histories, error) {
	data, err := r.repo.Update(data, params)
	if data != nil {
		fmt.Println("Data Terupdate")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
