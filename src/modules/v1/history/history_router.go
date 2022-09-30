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

	route.HandleFunc("/", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/save", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Add)).Methods("POST")
	route.HandleFunc("/delete/{history_id}", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("/update/{history_id}", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Update)).Methods("PUT")

	//sort and search
	route.HandleFunc("/sort/{params}", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Sort)).Methods("GET")
	route.HandleFunc("/search/{params}", middleware.HandlerChain(middleware.CheckAuth("user")).Then(ctrl.Search)).Methods("GET")

}
