package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloTableTest(t *testing.T) {
	tes := []struct {
		nama       string
		permintaan string
		ekspektasi string
	}{
		{
			nama:       "HelloAsep",
			permintaan: "Asep",
			ekspektasi: "Hello Asep",
		},
		{
			nama:       "HelloPutra",
			permintaan: "Putra",
			ekspektasi: "Hello Putra",
		},
	}
	for _, v := range tes {
		t.Run(v.nama, func(t *testing.T) {
			result := Hello(v.permintaan)
			assert.Equal(t, v.ekspektasi, result)
		})
	}
}
