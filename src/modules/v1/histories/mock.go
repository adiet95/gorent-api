package histories

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) HisFindAll() (*models.Histories, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Histories), nil
}

func (m *RepoMock) HisSave(data *models.History) (*models.History, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.History), nil
}
func (m *RepoMock) HisUpdate(data *models.History, id int) (*models.History, error) {
	args := m.mock.Called(data, id)
	return args.Get(0).(*models.History), nil
}
func (m *RepoMock) HisDelete(id int) (*models.History, error) {
	args := m.mock.Called(id)
	return args.Get(0).(*models.History), nil
}
func (m *RepoMock) FindByName(name string) (*models.Histories, error) {
	args := m.mock.Called(name)
	return args.Get(0).(*models.Histories), nil
}
func (m *RepoMock) GetFavo() (*models.Histories, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Histories), nil
}
