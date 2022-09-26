package history

import (
	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/save", middleware.CheckAuth(ctrl.Add)).Methods("POST")
	route.HandleFunc("/delete/{history_id}", middleware.CheckAuth(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("/update/{history_id}", middleware.CheckAuth(ctrl.Update)).Methods("PUT")

	//sort and search
	route.HandleFunc("/sort/{params}", ctrl.Sort).Methods("GET")
	route.HandleFunc("/search/{params}", ctrl.Search).Methods("GET")

}
