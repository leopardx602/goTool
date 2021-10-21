// go test -v -bench=. -run=none .

package unitTest_benchmark

import "testing"

// add "Benchmark" before function name
func BenchmarkGenByAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenByAppend()
	}
}

func BenchmarkGenByAppendCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenByAppendCap()
	}
}

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenByAssign()
	}
}
