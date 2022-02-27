package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(1, 2), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
