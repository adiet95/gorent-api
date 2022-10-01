package histories

import (
	"github.com/adiet95/gorent-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewHis(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()
	repo := NewRepoHis(db)
	svc := NewServiceHis(repo)
	ctrl := NewCtrlHis(svc)

	route.HandleFunc("/", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/", middleware.CheckAuth(ctrl.Add)).Methods("POST")
	route.HandleFunc("/", middleware.CheckAuth(ctrl.Update)).Methods("PUT")
	route.HandleFunc("/", middleware.CheckAuth(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("/search", middleware.CheckAuth(ctrl.Search)).Methods("POST")
	route.HandleFunc("/favo", middleware.CheckAuth(ctrl.Favorite)).Methods("GET")

}
