package interfaces

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
)

type VehiRepo interface {
	VeFindAll() (*models.Vehicles, error)
	VeSave(data *models.Vehicle) (*models.Vehicle, error)
	VeUpdate(data *models.Vehicle, id int) (*models.Vehicle, error)
	VeDelete(id int) (*models.Vehicle, error)
	FindByName(name string) (*models.Vehicles, error)
	PopularInCity(city string) (*models.Vehicles, error)
}

type VehiService interface {
	GetAll() *helpers.Response
	Add(data *models.Vehicle) *helpers.Response
	Update(data *models.Vehicle, id int) *helpers.Response
	Delete(id int) *helpers.Response
	Search(name string) *helpers.Response
	Popular(city string) *helpers.Response
}
