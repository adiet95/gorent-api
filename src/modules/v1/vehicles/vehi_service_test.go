package vehicles

import (
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceVe(&repo)

	var dataMocks = models.Vehicles{
		{Vehicle_Name: "Mobil Langkah", VehicleId: 1},
		{Vehicle_Name: "Motor Sport", VehicleId: 2},
	}

	repo.mock.On("VeFindAll").Return(&dataMocks, nil)

	data := service.GetAll()
	res := data.Data.(*models.Vehicles)

	for i, v := range *res {
		assert.Equal(t, dataMocks[i].VehicleId, v.VehicleId)
	}

}

func TestAdd(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceVe(&repo)

	var dataMock = models.Vehicle{Vehicle_Name: "Mobil Langkah", Type: "Mobil"}

	repo.mock.On("VeSave", &dataMock).Return(&dataMock, nil)
	data := service.Add(&dataMock)

	res := data.Data.(*models.Vehicle)
	var expect string = "Mobil"

	assert.Equal(t, expect, res.Type, "Vehicle name must be kuda")
}

func TestUpdate(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceVe(&repo)

	var dataMock = models.Vehicle{Vehicle_Name: "Kuda Sport", Type: "Kuda"}

	repo.mock.On("VeUpdate", &dataMock, 1).Return(&dataMock, nil)

	data := service.Update(&dataMock, 1)
	res := data.Data.(*models.Vehicle)
	expect := "Kuda Sport"

	assert.Equal(t, expect, res.Vehicle_Name, "Favorite must be updated")
}

func TestDelete(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceVe(&repo)

	var dataMock = models.Vehicle{Vehicle_Name: "Motor Sport", VehicleId: 2}

	repo.mock.On("VeDelete", 2).Return(&dataMock, nil)

	data := service.Delete(2)
	res := data

	assert.Equal(t, 200, res.Code, "Status code must be 200")
}
