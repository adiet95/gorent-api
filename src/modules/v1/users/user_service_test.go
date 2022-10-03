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

func TestUpdate(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{Email: "admin", Role: "admin"}
	var dataMocks = models.User{Email: "admin", Role: "admin", Address: "Jakarta", Phone: "+62813"}

	repo.mock.On("FindByEmail", "admin").Return(&dataMock, nil)
	repo.mock.On("UpdateUser", &dataMocks, "admin").Return(&dataMocks, nil)

	data := service.Update(&dataMocks, "admin", "image.jpg", "./uploads/")
	res := data.Data.(*models.User)
	var expect string = "Jakarta"

	assert.Equal(t, expect, res.Address, "Address must be Updated")
}

func TestDelete(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{Email: "user", Role: "user"}

	repo.mock.On("DeleteUser", "user").Return(&dataMock, nil)

	data := service.Delete("user")
	res := data

	assert.Equal(t, 204, res.Code, "Status code must be 204")
}
