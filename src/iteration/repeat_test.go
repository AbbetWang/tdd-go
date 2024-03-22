package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("can repeta 'a' five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("can repeta 'b' five times", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
