package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	var prev rune

	for _, r := range s {
		if (prev < 1 || unicode.IsDigit(prev)) && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if prev < 1 || unicode.IsDigit(prev) {
			prev = r
			continue
		}

		if !unicode.IsDigit(r) {
			builder.WriteRune(prev)
		} else {
			count := int(r) - '0'
			if count > 0 {
				builder.WriteString(strings.Repeat(string(prev), count))
			}
		}
		prev = r
	}

	if prev > 0 && !unicode.IsDigit(prev) {
		builder.WriteRune(prev)
	}

	return builder.String(), nil
}
