package pkg

import "unicode"

func upperFirstRune(r []rune) string {
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

func Title(s string) string {
	return upperFirstRune([]rune(s))
}
