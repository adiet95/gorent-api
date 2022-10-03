package histories

import (
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceHis(&repo)

	var dataMocks = models.Histories{
		{HistoryName: "Menyewa Motor", HistoryId: 1},
		{HistoryName: "Menyewa Mobil", HistoryId: 2},
	}

	repo.mock.On("HisFindAll").Return(&dataMocks, nil)

	data := service.GetAll()
	res := data.Data.(*models.Histories)

	for i, v := range *res {
		assert.Equal(t, dataMocks[i].HistoryId, v.HistoryId)
	}

}

func TestAdd(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceHis(&repo)

	var dataMock = models.History{HistoryName: "Menyewa Kuda", Favorite: true}

	repo.mock.On("HisSave", &dataMock).Return(&dataMock, nil)
	data := service.Add(&dataMock)

	res := data.Data.(*models.History)
	var expect string = "Menyewa Kuda"

	assert.Equal(t, expect, res.HistoryName, "History name must be kuda")
}

func TestUpdate(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceHis(&repo)

	var dataMock = models.History{HistoryName: "Menyewa Kuda", Favorite: true}

	repo.mock.On("HisUpdate", &dataMock, 1).Return(&dataMock, nil)

	data := service.Update(&dataMock, 1)
	res := data.Data.(*models.History)
	expect := "Menyewa Kuda"

	assert.Equal(t, expect, res.HistoryName, "Favorite must be updated")
}

func TestDelete(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewServiceHis(&repo)

	var dataMock = models.History{HistoryName: "Menyewa Mobil", HistoryId: 2}

	repo.mock.On("HisDelete", 2).Return(&dataMock, nil)

	data := service.Delete(2)
	res := data

	assert.Equal(t, 200, res.Code, "Status code must be 200")
}
