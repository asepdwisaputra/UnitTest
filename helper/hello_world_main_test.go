package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Aplikasi Dimulai")
	m.Run() // Eksekusi semua UnitTest
	fmt.Println("Aplikasi Selesai")
}
