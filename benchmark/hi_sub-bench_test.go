package benchmark

import (
	"testing"
)

func BenchmarkHiSub(b *testing.B) {
	b.Run("Asep", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hi("Asep")
		}
	})
	b.Run("Putra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hi("Putra")
		}
	})
}
