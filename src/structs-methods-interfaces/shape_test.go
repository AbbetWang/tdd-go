package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		t.Helper()
		got := Perimeter(10.0, 10.0)
		want := 40.0
		checkRectangles(got, want, t)
		got = Perimeter(15.0, 10.0)
		want = 50.0
		checkRectangles(got, want, t)
	})
}

func checkRectangles(got float64, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
