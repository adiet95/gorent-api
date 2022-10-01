package histories

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrlHis(reps interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{svc: reps}
}

func (re *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	re.svc.GetAll().Send(w)
}

func (re *history_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.History
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Add(&data).Send(w)
}

func (re *history_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	var datas models.History
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Update(&datas, v).Send(w)
}

func (re *history_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("id")
	v, err := strconv.Atoi(val)

	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Delete(v).Send(w)
}

func (re *history_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(w)
}

func (re *history_ctrl) Favorite(w http.ResponseWriter, r *http.Request) {
	re.svc.Favorite().Send(w)
}
