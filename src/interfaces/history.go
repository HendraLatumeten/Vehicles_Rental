package interfaces

import "github.com/hendralatumeten/vehicles_rental/src/database/orm/models"

type HistoryRepo interface {
	FindAll() (*models.HistoriesAll, error)
	Save(data *models.Histories) (*models.Histories, error)
	Delete(data *models.Histories, params string) (*models.Histories, error)
	Update(data *models.Histories, params string) (*models.Histories, error)
	Sort(params string) (*models.HistoriesAll, error)
	Search(params string) (*models.Histories, error)
}

type HistoryService interface {
	GetAll() (*models.HistoriesAll, error)
	Add(data *models.Histories) (*models.Histories, error)
	DeleteData(data *models.Histories, params string) (*models.Histories, error)
	UpdateData(data *models.Histories, params string) (*models.Histories, error)
	SortData(parmas string) (*models.HistoriesAll, error)
	SearchData(parmas string) (*models.Histories, error)
}
