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

// func (r *users_service) Add(data *models.User) *libs.Response {

// 	if check := r.repo.UserExsist(data.Username); check {
// 		return libs.Respone("data", 200, true)
// 	}
// 	r.repo.Save(data)
// 	//data, err := r.repo.Save(data)
// 	// if data != nil {
// 	// 	fmt.Println("Data Berhasil Disimpan")
// 	// }
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return libs.Respone("data", 200, true)
// }
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
func (r *users_service) UpdateData(data *models.User, params string) (*models.User, error) {
	data, err := r.repo.Update(data, params)
	if data != nil {
		fmt.Println("Data Terupdate")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *users_service) GetByUsername(username string) *libs.Response {
	data, err := r.repo.FindByUsername(username)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(data, 200, false)
}

func (r *users_service) Add(data *models.User) *libs.Response {

	//fmt.Println(data.Username)
	if check := r.repo.UserExsist(data.Username, data.Email); check {
		fmt.Println("data suda ada")
		return libs.Respone("username sudah terdaftar", 400, true)
	}

	hassPaasword, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}
	data.Password = hassPaasword
	result, err := r.repo.Save(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(result, 200, false)
}
