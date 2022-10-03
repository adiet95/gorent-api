package vehicles

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) VeFindAll() (*models.Vehicles, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Vehicles), nil
}
func (m *RepoMock) VeSave(data *models.Vehicle) (*models.Vehicle, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Vehicle), nil
}
func (m *RepoMock) VeUpdate(data *models.Vehicle, id int) (*models.Vehicle, error) {
	args := m.mock.Called(data, id)
	return args.Get(0).(*models.Vehicle), nil
}
func (m *RepoMock) VeDelete(id int) (*models.Vehicle, error) {
	args := m.mock.Called(id)
	return args.Get(0).(*models.Vehicle), nil
}
func (m *RepoMock) FindByName(name string) (*models.Vehicles, error) {
	args := m.mock.Called(name)
	return args.Get(0).(*models.Vehicles), nil
}
func (m *RepoMock) PopularInCity(city string) (*models.Vehicles, error) {
	args := m.mock.Called(city)
	return args.Get(0).(*models.Vehicles), nil
}
