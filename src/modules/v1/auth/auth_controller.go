package auth

import (
	"encoding/json"
	"net/http"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type auth_ctrl struct {
	rep interfaces.AuthService
}

func NewCtrl(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (a *auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}
	a.rep.Login(data).Send(w)
}
func (a *auth_ctrl) SignUp(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
	}

	a.rep.Register(data).Send(w)
}
