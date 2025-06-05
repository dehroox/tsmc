package main

import (
	"strings"
)

func LexData(src string) []string {
	nonWhitespaced := strings.Fields(src)
	final := make([]string, 0)

	for i := range nonWhitespaced {
		switch nonWhitespaced[i] {
		case "const":
			final = append(final, Token{Const, Keyword, "", ""}.ToString())

		}

		print(i)
	}

	return final
}
