package iteration

func Repeat(item string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += item
	}
	return repeated
}
