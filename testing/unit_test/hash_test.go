package unittest

import (
	"reflect"
	"testing"

	"github.com/adiet95/gorent-api/src/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	res, err := helpers.HashPassword("admin")
	temp := reflect.TypeOf(res)
	if err != nil {
		t.Fatal("error unit test")
	}
	assert.Equal(t, "string", temp)
}
