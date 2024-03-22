package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	println(strings.Count("cheese", "e"))
	t.Run("can repeta 'a' expected times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
