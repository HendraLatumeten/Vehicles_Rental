package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//method ke controller
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/save", ctrl.Add).Methods("POST")

}
