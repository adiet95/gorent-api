package vehicles

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type vehi_service struct {
	vehi_repo interfaces.VehiRepo
}

func NewServiceVe(reps interfaces.VehiRepo) *vehi_service {
	return &vehi_service{reps}
}

func (r *vehi_service) GetAll() *helpers.Response {
	data, err := r.vehi_repo.VeFindAll()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *vehi_service) Add(data *models.Vehicle) *helpers.Response {
	result, err := re.vehi_repo.VeSave(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 200, false)
}

func (re *vehi_service) Update(data *models.Vehicle, id int) *helpers.Response {
	data, err := re.vehi_repo.VeUpdate(data, id)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *vehi_service) Delete(id int) *helpers.Response {
	data, err := re.vehi_repo.VeDelete(id)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *vehi_service) Search(name string) *helpers.Response {
	data, err := re.vehi_repo.FindByName(name)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *vehi_service) Popular(city string) *helpers.Response {
	data, err := re.vehi_repo.PopularInCity(city)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}
