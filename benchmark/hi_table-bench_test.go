package benchmark

import "testing"

func BenchmarkHiTable(b *testing.B) {
	benchmarks := []struct {
		nama    string
		request string
	}{
		{
			nama:    "HiDwi",
			request: "Dwi",
		},
		{
			nama:    "HiPutra",
			request: "Putra",
		},
	}
	for _, bench := range benchmarks {
		b.Run(bench.nama, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hi(bench.request)
			}
		})
	}
}
