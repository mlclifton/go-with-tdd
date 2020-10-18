package iteration

// CharRepeater returns a string of character `char` repeated `n` times
func CharRepeater(char string, howMany int) string {
	res := char
	for n := 1; n < howMany; n++ {
		res += char
	}
	return res
}
