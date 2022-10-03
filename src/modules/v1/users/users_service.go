package users

//menghandle Logic
import (
	"fmt"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type users_service struct {
	repo interfaces.UsersRepo
}

func NewService(repo interfaces.UsersRepo) *users_service {
	return &users_service{repo}
}

func (r *users_service) GetAll() *libs.Response {
	data, err := r.repo.FindAll()

	if err != nil {
		return libs.Respone(err, 400, true)
	}
	return libs.Respone(data, 200, false)
}

func (r *users_service) GetByUsername(username string) *libs.Response {
	data, err := r.repo.FindByUsername(username)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(data, 200, false)
}

func (r *users_service) Add(data *models.User, filename string) *libs.Response {

	if check := r.repo.UserExsist(data.Username, data.Email); check {
		return libs.Respone("username atau email sudah terdaftar", 400, true)
	}

	hassPaasword, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	data.Image = filename
	data.Password = hassPaasword

	result, err := r.repo.Save(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(result, 200, false)
}

func (r *users_service) DeleteData(data *models.User, params string) (*models.User, error) {
	data, err := r.repo.Delete(data, params)
	if data != nil {
		fmt.Println("Data Terhapus")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *users_service) UpdateData(data *models.User, params string, filename string) *libs.Response {
	result, err := r.repo.Update(data, params, filename)
	if data != nil {
		fmt.Println("Data Terupdate")
	}

	if err != nil {
		return nil
	}
	return libs.Respone(result, 200, false)
}
