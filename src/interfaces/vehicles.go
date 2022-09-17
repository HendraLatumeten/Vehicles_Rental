package interfaces

import "github.com/hendralatumeten/vehicles_rental/src/database/orm/models"

type VehiclesRepo interface {
	FindAll() (*models.Vehicles, error)
	Save(data *models.Vehicle) (*models.Vehicle, error)
	Delete(data *models.Vehicle, params string) (*models.Vehicle, error)
	Update(data *models.Vehicle, params string) (*models.Vehicle, error)
}

type VehiclesService interface {
	GetAll() (*models.Vehicles, error)
	Add(data *models.Vehicle) (*models.Vehicle, error)
	DeleteData(data *models.Vehicle, params string) (*models.Vehicle, error)
	UpdateData(data *models.Vehicle, params string) (*models.Vehicle, error)
}
