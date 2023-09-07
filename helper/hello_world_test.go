package helper

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Asep")
	if result != "Hello Asep" {
		// Test Eror
		panic("Program Eror")
	}
}

// func TestHelloGuys(t *testing.T) {
// 	result := Hello("Semua")
// 	if result != "Hello Aja" { //Disini kita jadikan eror
// 		// Test Eror
// 		t.Error("Program Eror") // Jika eror akan dicetak pesannya
// 	}
// }
