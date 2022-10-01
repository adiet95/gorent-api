package histories

import (
	"errors"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepoHis(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (r *history_repo) HisFindAll() (*models.Histories, error) {
	var data models.Histories

	result := r.db.Order("history_id asc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &data, nil
}

func (r *history_repo) HisSave(data *models.History) (*models.History, error) {
	// var datas models.Histories
	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}

func (re *history_repo) HisUpdate(data *models.History, id int) (*models.History, error) {

	res := re.db.Model(&data).Where("history_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *history_repo) HisDelete(id int) (*models.History, error) {
	var data *models.History
	var datas *models.Histories
	res := re.db.Where("history_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("history_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *history_repo) FindByName(name string) (*models.Histories, error) {
	var datas *models.Histories
	res := re.db.Order("history_id asc").Where("LOWER(history_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *history_repo) GetFavo() (*models.Histories, error) {
	var datas models.Histories

	result := r.db.Order("history_id asc").Where("favorite = ?", true).Find(&datas)

	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &datas, nil
}
