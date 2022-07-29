package sequence

import "testing"

func BenchmarkGetResult(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetResult()
    }
}