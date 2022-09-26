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

	//route.HandleFunc("/save", middleware.MultipleMiddleware(middleware.CheckAuth(middleware.FileUpload(ctrl.GetAll)))).Methods("GET")
	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/save", middleware.FileUpload(ctrl.Add)).Methods("POST")
	route.HandleFunc("/delete/{user_id}", middleware.CheckAuth(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("/update/{user_id}", middleware.CheckAuth(middleware.FileUpload(ctrl.Update))).Methods("PUT")

}
