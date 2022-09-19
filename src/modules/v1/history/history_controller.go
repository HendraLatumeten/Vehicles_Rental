package history

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/responses"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(reps interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: reps}
}

func (re *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := re.svc.GetAll()
	if err != nil {
		//error handling
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		//response
		responses.JSON(w, http.StatusOK, &data)
	}
}

func (re *history_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	var decode = schema.NewDecoder()
	var datas models.Histories
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

func (re *history_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	historyId := mux.Vars(r)["history_id"]
	var datas models.Histories

	json.NewDecoder(r.Body).Decode(&datas)
	data, err := re.svc.DeleteData(&datas, historyId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}
func (re *history_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	historyId := mux.Vars(r)["history_id"]
	var decode = schema.NewDecoder()
	var datas models.Histories

	err := r.ParseForm()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	decode.Decode(&datas, r.PostForm)
	data, err := re.svc.UpdateData(&datas, historyId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusCreated, &data)
	}
}

//sort and search
func (re *history_ctrl) Sort(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SortData(params)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}
func (re *history_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["params"]
	data, err := re.svc.SearchData(params)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}
}
