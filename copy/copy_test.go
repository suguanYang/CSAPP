package copy

import (
	"testing"
)

type args struct {
	src [2048][2048]int
	dst [2048][2048]int
}

var tt args = args{src: GenerateTestDataSrc()}

func BenchmarkCopyij(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Copyij(&tt.src, &tt.dst)
	}
}

func BenchmarkCopyji(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Copyji(&tt.src, &tt.dst)
	}
}
