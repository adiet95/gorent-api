package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/adiet95/gorent-api/src/helpers"
)

func CheckAuthor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New("invalid header type", 401, true).Send(w)
			return
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.New(err.Error(), 401, true).Send(w)
			return
		}
		if checkToken.Role != "admin" {
			helpers.New("forbidden access", 401, true).Send(w)
			return
		}
		ctx := context.WithValue(r.Context(), "role", checkToken.Role)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
