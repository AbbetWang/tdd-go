package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		want := 3
		assertCounter(t, counter, want)
	})
}

func assertCounter(t testing.TB, counter Counter, want int) {
	t.Helper()
	if counter.Value() != 3 {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
