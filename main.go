package helpers

func Reverse(str string) (string, error) {
	reversed := ""

	for _, character := range str {
		reversed = string(character) + reversed
	}

	return reversed, nil
}
