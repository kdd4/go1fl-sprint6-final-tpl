package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseTranslator(text string) string {
	isNotMorseRune := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}

	if strings.ContainsFunc(text, isNotMorseRune) {
		return morse.ToMorse(text)
	}
	return morse.ToText(text)
}
