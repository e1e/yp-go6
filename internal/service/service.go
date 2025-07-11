package service

import (
	"strings"
	"yp-go6/pkg/morse"
)

func Convert(str string) string {
	convertString := morse.ToMorse(morse.ToText(str))

	if strings.Compare(str, convertString) == 0 {
		return morse.ToText(str)
	} else {
		return morse.ToMorse(str)
	}
}
