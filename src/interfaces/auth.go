package interfaces

import (
	"net/http"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
)

type AuthService interface {
	Login(body models.User, w http.ResponseWriter) *helpers.Response
	Register(body *models.User) *helpers.Response
}
