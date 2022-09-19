package vehicles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/responses"
)

// berinterakti dengan service dan router
// untuk mengahandle request yang masuk

type vehicles_ctrl struct {
	svc interfaces.VehiclesService
}

func NewCtrl(reps interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{svc: reps}
}

func (re *vehicles_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.GetAll()
	if err != nil {
		//error handling
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		//response
		responses.JSON(w, http.StatusOK, &data)
	}
}

func (re *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	var decode = schema.NewDecoder()
	var datas models.Vehicle
	err := r.ParseForm()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	decode.Decode(&datas, r.PostForm)

	data, err := re.svc.Add(&datas)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusCreated, &data)
	}
}

func (re *vehicles_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vehiclesId := mux.Vars(r)["vehicles_id"]
	var datas models.Vehicle

	json.NewDecoder(r.Body).Decode(&datas)
	data, err := re.svc.DeleteData(&datas, vehiclesId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}

func (re *vehicles_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	vehiclesId := mux.Vars(r)["vehicles_id"]
	var decode = schema.NewDecoder()
	var datas models.Vehicle

	err := r.ParseForm()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	decode.Decode(&datas, r.PostForm)
	data, err := re.svc.UpdateData(&datas, vehiclesId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusCreated, &data)
	}
}

//sort and search
func (re *vehicles_ctrl) Sort(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SortData(params)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}
func (re *vehicles_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SearchData(params)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}

//popular vehicles
func (re *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.PopularVehicles()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusCreated, &data)
	}
}
