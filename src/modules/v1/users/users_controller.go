package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/responses"
)

//menghandle Request
type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc: reps}
}

func (re *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := re.svc.GetAll()
	if err != nil {
		//error handling
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		//response
		responses.JSON(w, http.StatusOK, &data)
	}

}

func (re *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	var decode = schema.NewDecoder()
	var datas models.User
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
func (re *users_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userId := mux.Vars(r)["user_id"]
	var datas models.User

	json.NewDecoder(r.Body).Decode(&datas)
	data, err := re.svc.DeleteData(&datas, userId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	} else {
		responses.JSON(w, http.StatusOK, &data)
	}

}
func (re *users_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")

	userId := mux.Vars(r)["user_id"]
	var decode = schema.NewDecoder()
	var datas models.User

	err := r.ParseForm()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	decode.Decode(&datas, r.PostForm)
	data, err := re.svc.UpdateData(&datas, userId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	} else {
		responses.JSON(w, http.StatusCreated, &data)
	}

}
