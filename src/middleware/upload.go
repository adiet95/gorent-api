package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/adiet95/gorent-api/src/helpers"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const max = 1024 * 1024 // 1MB
		r.Body = http.MaxBytesReader(w, r.Body, max)
		if err := r.ParseMultipartForm(max); err != nil {
			helpers.New("The uploaded file is less than 1MB in size", 400, true)
			return
		}
		// check attribut file yang di upload
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			helpers.New("invalid attribute", 401, true).Send(w)
			return
		}
		defer file.Close()

		//Checking extension
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			helpers.New("Extension file not allowed", 401, true).Send(w)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Membuat folder file upload
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			helpers.New("error build file location", 401, true).Send(w)
			return
		}

		// Membuat file baru di direktori file upload
		fileName := r.FormValue("file_name")
		temp := fmt.Sprintf("%d-%s-%s", time.Now().UnixNano(), fileName, filepath.Ext(fileHeader.Filename))
		res := fmt.Sprintf("./uploads/%d-%s-%s", time.Now().UnixNano(), fileName, filepath.Ext(fileHeader.Filename))
		dst, err := os.Create(res)
		if err != nil {
			helpers.New("error while upload file", 401, true).Send(w)
			return
		}
		var name interface{} = temp

		defer dst.Close()

		// Mengcopy file ke filesistem sesuai direktorinya
		_, err = io.Copy(dst, file)
		if err != nil {
			helpers.New("error copy filesystem", 401, true).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "dir", name)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
