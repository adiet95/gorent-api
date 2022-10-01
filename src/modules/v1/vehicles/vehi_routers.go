package vehicles

import (
	"github.com/adiet95/gorent-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewVe(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehi").Subrouter()
	repo := NewRepoVe(db)
	svc := NewServiceVe(repo)
	ctrl := NewCtrlVe(svc)

	route.HandleFunc("/", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Add))).Methods("POST")
	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Update))).Methods("PUT")
	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Delete))).Methods("DELETE")
	route.HandleFunc("/search", middleware.CheckAuth(ctrl.Search)).Methods("POST")
	route.HandleFunc("/popular", middleware.CheckAuth(ctrl.Popular)).Methods("POST")

}
