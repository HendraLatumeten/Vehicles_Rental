package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/middleware"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)
	route.HandleFunc("/", ctrl.GetAll).Methods("GET")

	route.HandleFunc("/save", middleware.HandlerChain(middleware.CheckAuth("admin"), middleware.FileUpload).Then(ctrl.Add)).Methods("POST")
	route.HandleFunc("/delete/{vehicles_id}", middleware.HandlerChain(middleware.CheckAuth("admin")).Then(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("/update/{vehicles_id}", middleware.HandlerChain(middleware.CheckAuth("admin"), middleware.FileUpload).Then(ctrl.Update)).Methods("PUT")
	//sort and search

	route.HandleFunc("/sort/{params}", ctrl.Sort).Methods("GET")
	route.HandleFunc("/search/{params}", ctrl.Search).Methods("GET")
	//popular vehicles
	route.HandleFunc("/popular", ctrl.PopularVehicles).Methods("GET")

}
