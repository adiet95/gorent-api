package users

import (
	"net/http"

	"github.com/adiet95/gorent-api/src/database/orm/models"
	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/interfaces"
	"github.com/gorilla/schema"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")

	result := re.svc.FindEmail(claim_user.(string))
	result.Send(w)
}

func (re *user_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "multipart/form-data")
	var decode = schema.NewDecoder()
	var data models.User

	x := r.Context().Value("dir")
	res := x.(string)
	path := "./uploads/" + res
	data.FileName = res
	data.Path = path

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	decode.Decode(&data, r.Form)
	re.svc.Add(&data).Send(w)

}

func (re *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "multipart/form-data")
	claim_user := r.Context().Value("email")

	var decode = schema.NewDecoder()
	var data models.User

	x := r.Context().Value("dir")
	fileName := x.(string)
	path := "./uploads/" + fileName

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	decode.Decode(&data, r.Form)
	re.svc.Update(&data, claim_user.(string), fileName, path).Send(w)
}

func (re *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("email")
	re.svc.Delete(val).Send(w)
}
