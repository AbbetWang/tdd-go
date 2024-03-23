package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectoangle Rectangle) float64 {
	return 2 * (rectoangle.Width + rectoangle.Height)
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
