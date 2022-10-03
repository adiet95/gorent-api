package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/adiet95/gorent-api/src/middleware"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	res, err := helpers.HashPassword("admin")
	if err != nil {
		t.Fatal("error unit test")
	}
	assert.NotNil(t, res, "Must be not nil")
}

func TestNewToken(t *testing.T) {
	token := helpers.NewToken("admin", "admin")
	respon, err := token.Create()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, respon, "Must be not nil")
}

func TestCheckAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("email").(string); !ok {
			t.Errorf("email not in request context: got %q", val)
		}
	})

	rr := httptest.NewRecorder()

	handler := middleware.CheckAuth(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestCheckAuthor(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("role").(string); !ok {
			t.Errorf("email not in request context: got %q", val)
		}
	})

	rr := httptest.NewRecorder()

	handler := middleware.CheckAuthor(testHandler)
	handler.ServeHTTP(rr, req)
}

func TestUploadFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if val, ok := r.Context().Value("dir").(string); !ok {
			t.Errorf("email not in request context: got %q", val)
		}
	})

	rr := httptest.NewRecorder()

	handler := middleware.UploadFile(testHandler)
	handler.ServeHTTP(rr, req)
}
