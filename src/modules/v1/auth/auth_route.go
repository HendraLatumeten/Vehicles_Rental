package auth

import (
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/modules/v1/users"
	"gorm.io/gorm"
)

//method ke controller
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/login", ctrl.Signin).Methods("POST")
	route.HandleFunc("/register", ctrl.SignUp).Methods("POST")

}
