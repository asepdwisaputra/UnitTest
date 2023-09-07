package benchmark

import "testing"

func BenchmarkHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hi("Asep")
	}
}

func TestHi(t *testing.T) {
	result := Hi("Asep")
	if result != "Hi Asep" {
		t.Error("Program eror")
	}
}
