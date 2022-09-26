package interfaces

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type UsersRepo interface {
	FindAll() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	FindByRole(role string) (*models.User, error)
	UserExsist(username, email string) bool
	Save(data *models.User) (*models.User, error)

	Delete(data *models.User, params string) (*models.User, error)
	Update(data *models.User, params string, filename string) (*models.User, error)
	//Update(data *models.User, params string) (*models.User, error)
}

type UsersService interface {
	GetAll() (*models.Users, error)
	GetByUsername(username string) *libs.Response

	Add(data *models.User, filename string) *libs.Response
	DeleteData(data *models.User, params string) (*models.User, error)
	UpdateData(data *models.User, params string, filename string) (*models.User, error)
	//UpdateData(data *models.User, params string) (*models.User, error)
}
