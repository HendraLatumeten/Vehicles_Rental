package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
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
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)

}

func (re *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var decode = schema.NewDecoder()
	var datas models.Vehicle

	//convert context to string
	var x interface{} = r.Context().Value("image")
	filename := fmt.Sprintf("%v", x)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}
	decode.Decode(&datas, r.Form)
	re.svc.Add(&datas, filename).Send(w)

}

func (re *vehicles_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	vehiclesId := mux.Vars(r)["vehicles_id"]
	var datas models.Vehicle

	json.NewDecoder(r.Body).Decode(&datas)
	data, err := re.svc.DeleteData(&datas, vehiclesId)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)

}

func (re *vehicles_ctrl) Update(w http.ResponseWriter, r *http.Request) {

	vehiclesId := mux.Vars(r)["vehicles_id"]
	var decode = schema.NewDecoder()
	var datas models.Vehicle

	//convert context to string
	var x interface{} = r.Context().Value("image")
	filename := fmt.Sprintf("%v", x)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	decode.Decode(&datas, r.Form)
	data, err := re.svc.UpdateData(&datas, vehiclesId, filename)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)

}

//sort and search
func (re *vehicles_ctrl) Sort(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SortData(params)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	} else {
		libs.Respone(data, 200, false).Send(w)
	}
}
func (re *vehicles_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SearchData(params)
	if err != nil {
		libs.Respone(err, 400, true)
		return
	} else {
		libs.Respone(data, 200, false).Send(w)
	}
}

//popular vehicles
func (re *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.PopularVehicles()
	if err != nil {
		libs.Respone(err, 400, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)
}
