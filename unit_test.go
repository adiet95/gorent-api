package main

import (
	"testing"

	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	res, err := helpers.HashPassword("admin")
	if err != nil {
		t.Fatal("error unit test")
	}
	assert.NotNil(t, string(""), res)
}
