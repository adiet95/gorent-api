package histories

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type history_service struct {
	history_repo interfaces.HistoryRepo
}

func NewServiceHis(reps interfaces.HistoryRepo) *history_service {
	return &history_service{reps}
}

func (r *history_service) GetAll() *helpers.Response {
	data, err := r.history_repo.HisFindAll()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *history_service) Add(data *models.History) *helpers.Response {
	data, err := re.history_repo.HisSave(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *history_service) Update(data *models.History, id int) *helpers.Response {
	data, err := re.history_repo.HisUpdate(data, id)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *history_service) Delete(id int) *helpers.Response {
	data, err := re.history_repo.HisDelete(id)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *history_service) Search(name string) *helpers.Response {
	data, err := re.history_repo.FindByName(name)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *history_service) Favorite() *helpers.Response {
	data, err := re.history_repo.GetFavo()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}
