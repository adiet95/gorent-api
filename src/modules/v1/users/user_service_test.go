package users

import (
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindEmail(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{Email: "admin", Role: "admin"}
	var dataMocks = models.Users{
		{Email: "admin", Role: "admin"},
		{Email: "user", Role: "user"},
	}

	repo.mock.On("FindByEmail", "admin").Return(&dataMock, nil)
	repo.mock.On("FindAll").Return(&dataMocks, nil)

	data := service.FindEmail("admin")
	res := data.Data.(*models.Users)

	for i, v := range *res {
		assert.Equal(t, dataMocks[i].Email, v.Email)
	}

}

func TestAdd(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{Email: "user", Role: "user"}

	repo.mock.On("Save", &dataMock).Return(&dataMock, nil)
	data := service.Add(&dataMock)
	res := data.Data.(*models.User)
	var expectEmail string = "user"

	assert.Equal(t, expectEmail, res.Email, "Email must be User")
}
