package iteration

// Repeat returns the given character repeated n times.
func Repeat(character string, times int) (repeated string) {
	repeated = character
	for i := 1; i < times; i++ {
		repeated += character
	}
	return
}
