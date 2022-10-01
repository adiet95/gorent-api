package vehicles

import (
	"errors"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"gorm.io/gorm"
)

type vehi_repo struct {
	db *gorm.DB
}

func NewRepoVe(db *gorm.DB) *vehi_repo {
	return &vehi_repo{db}
}

func (r *vehi_repo) VeFindAll() (*models.Vehicles, error) {
	var datas *models.Vehicles
	result := r.db.Model(&datas).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *vehi_repo) VeSave(data *models.Vehicle) (*models.Vehicle, error) {
	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *vehi_repo) VeUpdate(data *models.Vehicle, id int) (*models.Vehicle, error) {
	res := re.db.Model(&data).Where("vehicle_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *vehi_repo) VeDelete(id int) (*models.Vehicle, error) {
	var data *models.Vehicle
	var datas *models.Vehicles
	res := re.db.Where("vehicle_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("vehicle_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *vehi_repo) FindByName(name string) (*models.Vehicles, error) {
	var datas *models.Vehicles
	res := re.db.Order("vehicle_id asc").Where("LOWER(vehicle_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *vehi_repo) PopularInCity(city string) (*models.Vehicles, error) {
	var datas *models.Vehicles
	res := re.db.Order("popular desc").Limit(3).Where("LOWER(city) LIKE ?", "%"+city+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}
