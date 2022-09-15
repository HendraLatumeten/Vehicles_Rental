package interfaces

import "github.com/hendralatumeten/vehicles_rental/src/database/orm/models"

type UsersRepo interface {
	FindAll() (*models.Users, error)
	Save(data *models.User) (*models.User, error)
}

type UsersService interface {
	GetAll() (*models.Users, error)
	Add(data *models.User) (*models.User, error)
}
