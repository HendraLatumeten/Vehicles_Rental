package auth

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/interfaces"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
)

type auth_service struct {
	repo interfaces.UsersRepo
}

type token_respone struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.UsersRepo) *auth_service {
	return &auth_service{reps}
}

func (a auth_service) Login(body models.User) *libs.Response {
	//get user database
	user, err := a.repo.FindByUsername(body.Username)
	if err != nil {
		return libs.Respone("username tidak terdaftar", 401, true)
	}

	//password check
	if !libs.CheckPassword(user.Password, body.Password) {
		return libs.Respone("Password salah", 401, true)
	}

	//create Token
	token := libs.NweToken(body.Username, user.Role)
	TokenCreate, err := token.Create()
	if err != nil {
		return libs.Respone(err.Error(), 401, true)
	}
	return libs.Respone(token_respone{Tokens: TokenCreate}, 200, false)
}
