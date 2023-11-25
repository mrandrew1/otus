package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

const escape = '\\'

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var out strings.Builder
	var c rune
	var flagEscape bool

	out.Grow(len(s))

	for index, value := range s {
		switch {
		case value == escape && !flagEscape:
			flagEscape = true
		case flagEscape:
			out.WriteRune(value)
			flagEscape, c = false, value
		case unicode.IsNumber(value):
			if c != 0 {
				for a := 1; a < int(value-'0'); a++ {
					out.WriteRune(c)
				}
				c = 0
			} else {
				return "", errors.New("invalid string")
			}
		default:
			if (len(s) > (index+1) && s[index+1] != '0') || (len(s) == (index + 1)) {
				out.WriteRune(value)
				c = value
			}
		}
	}
	return out.String(), nil
}
