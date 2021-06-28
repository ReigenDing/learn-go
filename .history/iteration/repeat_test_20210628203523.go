package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// 基准测试

func BenchMark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
