package users

import (
	"fmt"
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

func TestFindEmail(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Users{
		{Email: "admin", Role: "admin"},
		{Email: "user", Role: "user"},
	}

	repo.mock.On("FindAll").Return(&dataMock, nil)
	data := service.FindEmail("admin")
	result := data.Data.(*models.Users)
	fmt.Println(result)

}
