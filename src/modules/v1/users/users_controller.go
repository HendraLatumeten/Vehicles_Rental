package users

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

//menghandle Request
type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc: reps}
}

func (re *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	result := re.svc.GetByUsername(username.(string))
	result.Send(w)

}
func (re *users_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	var datas models.User

	json.NewDecoder(r.Body).Decode(&datas)
	data, err := re.svc.DeleteData(&datas, userId)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}
	libs.Respone(data, 200, false).Send(w)

}

func (re *users_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	var decode = schema.NewDecoder()
	var datas models.User

	//convert context to string
	var x interface{} = r.Context().Value("image")
	filename := fmt.Sprintf("%v", x)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}
	decode.Decode(&datas, r.Form)
	re.svc.UpdateData(&datas, userId, filename).Send(w)
}

func (re *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var decode = schema.NewDecoder()
	var datas models.User

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
