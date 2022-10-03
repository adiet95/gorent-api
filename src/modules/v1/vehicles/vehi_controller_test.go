package vehicles

import (
	"net/http/httptest"
	"testing"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = RepoMock{mock.Mock{}}
var service = NewServiceVe(&repo)
var ctrl = NewCtrlVe(service)

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMocks = models.Vehicles{
		{Vehicle_Name: "Mobil Langkah", VehicleId: 1},
		{Vehicle_Name: "Motor Sport", VehicleId: 2},
	}

	repo.mock.On("VeFindAll").Return(&dataMocks, nil)

	req := httptest.NewRequest("GET", "/test/history", nil)

	mux.HandleFunc("/test/history", ctrl.GetAll)
	mux.ServeHTTP(w, req)

	var Vehicle *models.Vehicle
	respon := helpers.Response{
		Data: &Vehicle,
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlAdd(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.Vehicle{Vehicle_Name: "Mobil Langkah", Type: "Mobil"}

	repo.mock.On("VeSave", &dataMock).Return(&dataMock, nil)
	req := httptest.NewRequest("POST", "/test/history", w.Body)

	mux.HandleFunc("/test/history", ctrl.Add)
	mux.ServeHTTP(w, req)

	var Vehicle *models.Vehicle
	respon := helpers.Response{
		Data: &Vehicle,
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.Vehicle{Vehicle_Name: "Kuda Sport", Type: "Kuda"}

	repo.mock.On("VeUpdate", &dataMock, 1).Return(&dataMock, nil)
	req := httptest.NewRequest("PUT", "/test/history", w.Body)

	mux.HandleFunc("/test/history/?id=2", ctrl.Update)
	mux.ServeHTTP(w, req)

	var Vehicle *models.Vehicle
	respon := helpers.Response{
		Data: &Vehicle,
	}
	// if err := json.Unmarshal(w.Body.Bytes(), &respon); err != nil {
	// 	t.Fatal(err)
	// }
	assert.Equal(t, 0, respon.Code, "status code must be 0")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlDelete(t *testing.T) {
	mux := mux.NewRouter()
	w := httptest.NewRecorder()

	var dataMock = models.Vehicle{Vehicle_Name: "Motor Sport", VehicleId: 2}
	repo.mock.On("VeDelete", 2).Return(&dataMock, nil)

	req := httptest.NewRequest("DELETE", "/test/history", nil)

	mux.HandleFunc("/test/history/?id=2", ctrl.Delete)
	mux.ServeHTTP(w, req)

	var Vehicle *models.Vehicle
	respon := helpers.Response{
		Data: &Vehicle,
	}
	assert.Equal(t, 0, respon.Code, "status code must be 204")
	assert.False(t, respon.IsError, "error must be false")
}
