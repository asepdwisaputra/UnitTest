package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWithTestify(t *testing.T) {
	// (t, ekspektasiHasil, realitaHasil) --> ekspektasi harus = realita
	assert.Equal(t, "Hello", "Hello")
}

// func TestHelloWithAssert(t *testing.T) {
// 	assert.Equal(t, "Hello", "Hai")
// 	fmt.Println("Program Selesai")
// }

// func TestHelloWithRequire(t *testing.T) {
// 	require.Equal(t, "Hello", "Hai")
// 	fmt.Println("Program Selesai") // Seharusnya ini tidak dijalankan
// }
