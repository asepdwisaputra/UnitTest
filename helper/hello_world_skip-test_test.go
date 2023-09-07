package helper

import "testing"

func TestHelloSkip(t *testing.T) {
	result := Hello("Hello2")
	if result == "Hello Hello2" {
		t.Skip("Tes ini dilewati")
	}
}
