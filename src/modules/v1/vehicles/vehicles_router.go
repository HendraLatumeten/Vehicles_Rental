package vehicles

import (
	"github.com/gorilla/mux"
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
	route.HandleFunc("/save", ctrl.Add).Methods("POST")
	route.HandleFunc("/delete/{vehicles_id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update/{vehicles_id}", ctrl.Update).Methods("PUT")

}
