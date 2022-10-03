package histories

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
var service = NewServiceHis(&repo)
var ctrl = NewCtrlHis(service)

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMocks = models.Histories{
		{HistoryName: "Menyewa Motor", HistoryId: 1},
		{HistoryName: "Menyewa Mobil", HistoryId: 2},
	}

	repo.mock.On("HisFindAll").Return(&dataMocks, nil)

	req := httptest.NewRequest("GET", "/test/history", nil)

	mux.HandleFunc("/test/history", ctrl.GetAll)
	mux.ServeHTTP(w, req)

	var History *models.History
	respon := helpers.Response{
		Data: &History,
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlAdd(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.History{HistoryName: "Menyewa Kuda", Favorite: true}

	repo.mock.On("HisSave", &dataMock).Return(&dataMock, nil)
	req := httptest.NewRequest("POST", "/test/history", w.Body)

	mux.HandleFunc("/test/history", ctrl.Add)
	mux.ServeHTTP(w, req)

	var History *models.History
	respon := helpers.Response{
		Data: &History,
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.History{HistoryName: "Menyewa Kuda"}
	repo.mock.On("HisUpdate", &dataMock, 2).Return(&dataMock, nil)

	req := httptest.NewRequest("PUT", "/test/history", w.Body)

	mux.HandleFunc("/test/history/?id=2", ctrl.Update)
	mux.ServeHTTP(w, req)

	var History *models.History
	respon := helpers.Response{
		Data: &History,
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

	var dataMock = models.History{HistoryName: "Menyewa Mobil", HistoryId: 2}
	repo.mock.On("HisDelete", 2).Return(&dataMock, nil)

	req := httptest.NewRequest("DELETE", "/test/history", nil)

	mux.HandleFunc("/test/history/?id=2", ctrl.Delete)
	mux.ServeHTTP(w, req)

	var History *models.History
	respon := helpers.Response{
		Data: &History,
	}
	assert.Equal(t, 0, respon.Code, "status code must be 204")
	assert.False(t, respon.IsError, "error must be false")
}
