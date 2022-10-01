package interfaces

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
)

type HistoryRepo interface {
	HisFindAll() (*models.Histories, error)
	HisSave(data *models.History) (*models.History, error)
	HisUpdate(data *models.History, id int) (*models.History, error)
	HisDelete(id int) (*models.History, error)
	FindByName(name string) (*models.Histories, error)
	GetFavo() (*models.Histories, error)
}

type HistoryService interface {
	GetAll() *helpers.Response
	Add(data *models.History) *helpers.Response
	Update(data *models.History, id int) *helpers.Response
	Delete(id int) *helpers.Response
	Search(name string) *helpers.Response
	Favorite() *helpers.Response
}
