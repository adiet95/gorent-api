package auth

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type auth_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(reps interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{reps}
}

func (a *auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}

	a.repo.Login(data, w).Send(w)
}

func (a *auth_ctrl) Register(w http.ResponseWriter, r *http.Request) {
	var data *models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}
	a.repo.Register(data).Send(w)
}
