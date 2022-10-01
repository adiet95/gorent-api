package checkout

import (
	"errors"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"gorm.io/gorm"
)

type co_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *co_repo {
	return &co_repo{db}
}

func (r *co_repo) FindAll() (*models.Vehicleitems, error) {
	var datas *models.Vehicleitems
	result := r.db.Model(&datas).Preload("User").Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}

	return datas, nil
}

func (r *co_repo) FindData(id string) (*models.Vehicleitems, error) {
	var datas *models.Vehicleitems
	result := r.db.Model(&datas).Preload("User").Where("user_id = ?", id).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}

	return datas, nil
}

func (r *co_repo) Payment(data *models.Vehicleitem) error {
	var vehicle *models.Vehicle
	var vehicles *models.Vehicles

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&vehicles).Where("vehicle_id = ?", data.VehicleId).First(&vehicle).Error; err != nil {
		vehicle.Popular = vehicle.Popular + data.Amount
		r.db.Save(&vehicle)
		tx.Rollback()
		return err
	}

	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *co_repo) GetId(email string) (*models.User, error) {
	var users *models.Users
	var user *models.User

	result := r.db.Model(&users).Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, errors.New("invalid user_id")
	}
	return user, nil
}

func (r *co_repo) PaymentVehi(data *models.Vehicleitem) (*models.Vehicleitem, error) {
	var vehicle *models.Vehicle
	var vehicles *models.Vehicles

	result := r.db.Model(&vehicles).Where("vehicle_id = ?", data.VehicleId).First(&vehicle)
	if result.Error != nil {
		return nil, errors.New("invalid user_id")
	}
	vehicle.Popular = vehicle.Popular + data.Amount
	r.db.Save(&vehicle)

	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("invalid user_id")
	}

	return data, nil
}
