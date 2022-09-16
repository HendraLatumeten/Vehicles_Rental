package interfaces

import "github.com/hendralatumeten/vehicles_rental/src/database/orm/models"

type UsersRepo interface {
	FindAll() (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	Delete(data *models.User, params string) (*models.User, error)
	Update(data *models.User, params string) (*models.User, error)
}

type UsersService interface {
	GetAll() (*models.Users, error)
	Add(data *models.User) (*models.User, error)
	DeleteData(data *models.User, params string) (*models.User, error)
	UpdateData(data *models.User, params string) (*models.User, error)
}
