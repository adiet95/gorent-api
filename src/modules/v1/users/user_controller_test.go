package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = RepoMock{mock.Mock{}}
var service = NewService(&repo)
var ctrl = NewCtrl(service)

var dataMock = models.User{Email: "admin", Role: "admin"}
var dataMocks = models.Users{
	{Email: "admin", Role: "admin"},
	{Email: "user", Role: "user"},
}

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repo.mock.On("FindByEmail", "admin").Return(&dataMock, nil)
	repo.mock.On("FindAll").Return(&dataMocks, nil)

	req := httptest.NewRequest("GET", "/test/user", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "email", "admin")
	req = req.WithContext(ctx)

	mux.HandleFunc("/test/user", ctrl.GetAll)

	mux.ServeHTTP(w, req)

	var user *models.Users
	respon := helpers.Response{
		Data: &user,
	}

	fmt.Println(respon)

	if err := json.Unmarshal(w.Body.Bytes(), &respon); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlAdd(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.User{Email: "user", Role: "user"}

	repo.mock.On("Save", &dataMock).Return(&dataMock, nil)

	req := httptest.NewRequest("POST", "/test/user", w.Body)

	ctx := req.Context()
	ctx = context.WithValue(ctx, "dir", "./uploads/")
	req = req.WithContext(ctx)

	mux.HandleFunc("/test/user", ctrl.Add)

	mux.ServeHTTP(w, req)

	var user *models.User
	respon := helpers.Response{
		Data: &user,
	}

	fmt.Println(respon)

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}
