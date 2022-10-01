package checkout

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
)

type co_ctrl struct {
	svc interfaces.CoService
}

func NewCtrl(reps interfaces.CoService) *co_ctrl {
	return &co_ctrl{svc: reps}
}

func (re *co_ctrl) Checkout(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicleitem
	claim_user := r.Context().Value("email")
	email := claim_user.(string)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Checkout(&data, email).Send(w)
}

func (re *co_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")
	email := claim_user.(string)
	re.svc.GetAll(email).Send(w)
}
