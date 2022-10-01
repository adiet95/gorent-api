package auth

import (
	"github.com/adiet95/gorent-api/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewAu(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/").Subrouter()
	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.SignIn).Methods("POST")
	route.HandleFunc("/register", ctrl.Register).Methods("POST")

}
