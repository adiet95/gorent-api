package users

import (
	"os"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (re *user_service) Add(data *models.User) *helpers.Response {

	hassPass, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	data.Password = hassPass
	result, err := re.user_repo.Save(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 201, false)
}

func (re *user_service) Update(data *models.User, email string, fileName string, path string) *helpers.Response {
	//Get old path then remove the old file in directory
	oldData, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	oldPath := &oldData.Path
	if *oldPath != "" {
		err1 := os.Remove(*oldPath)
		if err1 != nil {
			return helpers.New(err1.Error(), 400, true)
		}
	}

	//Hasing New Password and update data
	hassPass, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	data.FileName = fileName
	data.Path = path
	data.Email = oldData.Email
	data.Role = oldData.Role
	data.Password = hassPass

	result, err := re.user_repo.UpdateUser(data, email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 202, false)
}

func (re *user_service) Delete(email string) *helpers.Response {

	data, err := re.user_repo.DeleteUser(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 204, false)
}

func (re *user_service) FindEmail(email string) *helpers.Response {
	data, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	if data.Role == "admin" {
		datas, err := re.user_repo.FindAll()
		if err != nil {
			return helpers.New(err.Error(), 400, true)
		}
		return helpers.New(datas, 200, false)
	}
	return helpers.New(data, 200, false)
}
