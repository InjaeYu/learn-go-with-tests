package iteration

import "testing"


func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a")
	}
}

func BenchmarkRepeat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat3("a")
	}
}