package users

import (
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/middleware"
	"gorm.io/gorm"
)

//method ke controller
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("user", "admin")).Then(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/save", middleware.HandlerChain(middleware.FileUpload).Then(ctrl.Add)).Methods("POST")
	route.HandleFunc("/delete/{user_id}", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Add)).Methods("DELETE")
	route.HandleFunc("/update/{user_id}", middleware.HandlerChain(middleware.CheckAuth("user"), middleware.FileUpload).Then(ctrl.Update)).Methods("PUT")

}
