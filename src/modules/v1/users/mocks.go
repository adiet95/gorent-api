package users

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindAll() (*models.Users, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Users), nil
}

func (m *RepoMock) Save(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) UpdateUser(data *models.User, email string) (*models.User, error) {
	args := m.mock.Called(data, email)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) DeleteUser(email string) (*models.User, error) {
	args := m.mock.Called(email)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) FindByEmail(email string) (*models.User, error) {
	args := m.mock.Called(email)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) RegisterEmail(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}
