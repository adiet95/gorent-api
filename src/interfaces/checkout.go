package interfaces

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
)

type CoRepo interface {
	FindAll() (*models.Vehicleitems, error)
	GetId(email string) (*models.User, error)
	Payment(data *models.Vehicleitem) error
	PaymentVehi(data *models.Vehicleitem) (*models.Vehicleitem, error)
	FindData(id string) (*models.Vehicleitems, error)
}

type CoService interface {
	GetAll(email string) *helpers.Response
	Checkout(data *models.Vehicleitem, email string) *helpers.Response
}
