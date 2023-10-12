package iteration

func Repeat(character string) string {
	result := character
	for i := 0; i < 4; i++ {
		result += character
	}
	return result
}
