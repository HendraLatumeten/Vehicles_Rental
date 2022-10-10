package interfaces

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type VehiclesRepo interface {
	FindAll() (*models.Vehicles, error)
	Save(data *models.Vehicle) (*models.Vehicle, error)
	Delete(data *models.Vehicle, params string) (*models.Vehicle, error)
	Update(data *models.Vehicle, params string, filename string) (*models.Vehicle, error)
	Sort(params string) (*models.Vehicles, error)
	Search(params string) (*models.Vehicles, error)
	Popular() (*models.Vehicles, error)
}

type VehiclesService interface {
	GetAll() (*models.Vehicles, error)
	Add(data *models.Vehicle) *libs.Response
	DeleteData(data *models.Vehicle, params string) (*models.Vehicle, error)
	UpdateData(data *models.Vehicle, params string, filename string) (*models.Vehicle, error)
	SortData(parmas string) (*models.Vehicles, error)
	SearchData(parmas string) (*models.Vehicles, error)
	PopularVehicles() (*models.Vehicles, error)
}

type UnitTest interface {
	GetAll() *models.Vehicles
}
