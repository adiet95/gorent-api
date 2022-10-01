package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type vehi_ctrl struct {
	svc interfaces.VehiService
}

func NewCtrlVe(reps interfaces.VehiService) *vehi_ctrl {
	return &vehi_ctrl{svc: reps}
}

func (re *vehi_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	re.svc.GetAll().Send(w)
}

func (re *vehi_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Add(&data).Send(w)
}

func (re *vehi_ctrl) Update(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Update(&datas, v).Send(w)
}

func (re *vehi_ctrl) Delete(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(w)
}

func (re *vehi_ctrl) Search(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("name")
	v := strings.ToLower(val)

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Search(v).Send(w)
}

func (re *vehi_ctrl) Popular(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("city")
	v := strings.ToLower(val)

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Popular(v).Send(w)
}
