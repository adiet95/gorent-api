package routers

import (
	"errors"

	"github.com/adiet95/gorent-api/src/database/orm"
	"github.com/adiet95/gorent-api/src/modules/v1/auth"
	"github.com/adiet95/gorent-api/src/modules/v1/checkout"
	"github.com/adiet95/gorent-api/src/modules/v1/histories"
	"github.com/adiet95/gorent-api/src/modules/v1/users"
	"github.com/adiet95/gorent-api/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	auth.NewAu(mainRoute, db)
	users.New(mainRoute, db)
	checkout.New(mainRoute, db)
	vehicles.NewVe(mainRoute, db)
	histories.NewHis(mainRoute, db)

	return mainRoute, nil
}
