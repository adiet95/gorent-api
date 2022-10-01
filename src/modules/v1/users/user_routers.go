package users

import (
	"github.com/adiet95/gorent-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckAuthor(middleware.UploadFile(ctrl.Add)))).Methods("POST")
	route.HandleFunc("/", middleware.CheckAuth(middleware.UploadFile(ctrl.Update))).Methods("PUT")
	route.HandleFunc("/", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Delete))).Methods("DELETE")
}
