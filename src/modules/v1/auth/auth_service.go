package auth

import (
	"net/http"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type auth_service struct {
	repo interfaces.UserRepo
}
type token_response struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.UserRepo) *auth_service {
	return &auth_service{reps}
}

func (a auth_service) Login(body models.User, w http.ResponseWriter) *helpers.Response {
	user, err := a.repo.FindByEmail(body.Email)
	if err != nil {
		return helpers.New("email not registered, register first", 401, true)
	}
	if !helpers.CheckPass(user.Password, body.Password) {
		return helpers.New("wrong password", 401, true)
	}
	token := helpers.NewToken(body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return helpers.New(err.Error(), 401, true)
	}

	w.Header().Set("Access", theToken)

	return helpers.New(token_response{Tokens: theToken}, 200, false)
}

func (a auth_service) Register(body *models.User) *helpers.Response {
	hassPass, err := helpers.HashPassword(body.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	body.Password = hassPass
	result, err := a.repo.RegisterEmail(body)
	if err != nil {
		return helpers.New("Email already registered, please login", 401, true)
	}
	return helpers.New(result, 200, false)
}
