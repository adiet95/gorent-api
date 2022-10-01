package interfaces

import (
	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	RegisterEmail(data *models.User) (*models.User, error)
}

type UserService interface {
	Add(data *models.User) *helpers.Response
	Update(data *models.User, email string, fileName string, path string) *helpers.Response
	Delete(email string) *helpers.Response
	FindEmail(email string) *helpers.Response
}
