package funcs

import "math/rand/v2"

func GeneratePass(size int, pass_type int) string {
	output := ""
	valid_chars := ""

	if pass_type >= 0 {
		valid_chars += CharsLetter
	}

	if pass_type >= 1 {
		valid_chars += CharsNum
	}

	if pass_type >= 2 {
		valid_chars += CharsSymbol
	}

	for range size {
		output += string(valid_chars[rand.IntN(len(valid_chars))])
	}

	return output
}
