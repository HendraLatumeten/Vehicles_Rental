package users

//menghandle Logic
import (
	"fmt"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
)

type users_service struct {
	repo interfaces.UsersRepo
}

func NewService(reps interfaces.UsersRepo) *users_service {
	return &users_service{reps}
}

func (r *users_service) GetAll() (*models.Users, error) {
	data, err := r.repo.FindAll()
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *users_service) Add(data *models.User) (*models.User, error) {
	data, err := r.repo.Save(data)
	if data != nil {
		fmt.Println("Data Berhasil Disimpan")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
