package history

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/save", ctrl.Add).Methods("POST")
	route.HandleFunc("/delete/{history_id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update/{history_id}", ctrl.Update).Methods("PUT")
	//sort and search
	route.HandleFunc("/sort/{params}", ctrl.Sort).Methods("GET")
	route.HandleFunc("/search/{params}", ctrl.Search).Methods("GET")

}
