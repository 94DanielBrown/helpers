package helpers

func Reverse(str string) string {
	reversed := ""

	for _, character := range str {
		reversed = string(character) + reversed
	}

	return reversed
}
