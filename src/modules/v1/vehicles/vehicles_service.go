package vehicles

//menghandle Logic
import (
	"fmt"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
)

type vehicles_service struct {
	repo interfaces.VehiclesRepo
}

func NewService(reps interfaces.VehiclesRepo) *vehicles_service {
	return &vehicles_service{reps}
}

func (r *vehicles_service) GetAll() (*models.Vehicles, error) {
	data, err := r.repo.FindAll()
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *vehicles_service) Add(data *models.Vehicle) (*models.Vehicle, error) {
	data, err := r.repo.Save(data)
	if data != nil {
		fmt.Println("Data Berhasil Disimpan")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (r *vehicles_service) DeleteData(data *models.Vehicle, params string) (*models.Vehicle, error) {
	data, err := r.repo.Delete(data, params)
	if data != nil {
		fmt.Println("Data Terhapus")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *vehicles_service) UpdateData(data *models.Vehicle, params string) (*models.Vehicle, error) {
	data, err := r.repo.Update(data, params)
	if data != nil {
		fmt.Println("Data Terupdate")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

//sort and search
func (r *vehicles_service) SortData(params string) (*models.Vehicles, error) {
	data, err := r.repo.Sort(params)
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *vehicles_service) SearchData(params string) (*models.Vehicles, error) {
	data, err := r.repo.Search(params)
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}

	if err != nil {
		return nil, err
	}
	return data, nil
}

//popular vehicles
func (r *vehicles_service) PopularVehicles() (*models.Vehicles, error) {
	data, err := r.repo.Popular()
	if data != nil {
		fmt.Println("Get Data Berhasil")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

// berinteraksi dengan repo dan controller
// berisi logic bisnis
