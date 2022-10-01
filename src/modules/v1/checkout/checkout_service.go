package checkout

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type co_service struct {
	co_repo interfaces.CoRepo
}

func NewService(reps interfaces.CoRepo) *co_service {
	return &co_service{reps}
}

func (re *co_service) Checkout(data *models.Vehicleitem, email string) *helpers.Response {
	res, err := re.co_repo.GetId(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	temp := res.UserId
	data.UserId = temp

	result, err := re.co_repo.PaymentVehi(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 200, false)
}

func (re *co_service) GetAll(email string) *helpers.Response {
	res, err := re.co_repo.GetId(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	if res.Role != "admin" {
		data, err := re.co_repo.FindData(res.UserId)
		if err != nil {
			return helpers.New(err.Error(), 400, true)
		}
		return helpers.New(data, 200, false)
	}

	data, err := re.co_repo.FindAll()
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}
