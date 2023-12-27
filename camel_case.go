package rkstrings

import (
	"strings"
	"unicode"
)

func ToCamelCaseLowercaseFirst(in string) string {
	return LowerCaseFirst(ToCamelCase(in))
}

func ToCamelCase(in string) string {
	var newName strings.Builder
	camelNext := true
	for _, elem := range in {
		if camelNext {
			newName.WriteRune(unicode.ToUpper(elem))
			camelNext = false
		} else if elem == '-' || elem == '_' {
			camelNext = true
		} else {
			newName.WriteRune(elem)
		}
	}
	return newName.String()
}

// backItUpOneBuilder wraps a strings.Builder that can back up one rune.
type backItUpOneBuilder struct {
	sb       strings.Builder
	lastRune rune
}

func (b *backItUpOneBuilder) WriteRune(r rune) {
	if b.lastRune != 0 {
		b.sb.WriteRune(b.lastRune)
	}
	b.lastRune = r
}

func (b *backItUpOneBuilder) BackItUp() {
	b.lastRune = 0
}

func (b *backItUpOneBuilder) String() string {
	b.sb.WriteRune(b.lastRune)
	return b.sb.String()
}

func ConvertCamelCase(in string, delim rune) string {
	var newName backItUpOneBuilder
	upperSequenceLen := 0
	var lastUpperRune rune
	for ii, elem := range in {
		if IsUpperRune(elem) {
			if ii != 0 && upperSequenceLen == 0 {
				newName.WriteRune(delim)
			}
			upperSequenceLen++
			lastUpperRune = elem
			newName.WriteRune(unicode.ToLower(elem))
		} else {
			if upperSequenceLen > 1 {
				newName.BackItUp()
				newName.WriteRune(delim)
				newName.WriteRune(unicode.ToLower(lastUpperRune))
			}
			upperSequenceLen = 0
			newName.WriteRune(elem)
		}
	}
	return newName.String()
}
