package checkout

import (
	"github.com/adiet95/gorent-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/checkout").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.CheckAuth(ctrl.Checkout)).Methods("POST")
	route.HandleFunc("/", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")

}
