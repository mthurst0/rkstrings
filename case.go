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

// convertFromCamelCase converts a camelCase string to a delimited string.
// The delimiter is specified by the delim parameter.
func convertFromCamelCase(s string, delim rune) string {
	var newName backItUpOneBuilder
	upperSequenceLen := 0
	var lastUpperRune rune
	for i, elem := range s {
		if IsUpperRune(elem) {
			if i != 0 && upperSequenceLen == 0 {
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

func ToSnakeCase(in string) string {
	return convertFromCamelCase(in, '_')
}

func IsLowerRune(r rune) bool {
	return unicode.IsLower(r) && unicode.IsLetter(r)
}

func IsUpperRune(r rune) bool {
	return unicode.IsUpper(r) && unicode.IsLetter(r)
}

func IsLowerCase(s string) bool {
	for _, c := range s {
		if !IsLowerRune(c) {
			return false
		}
	}
	return true
}

func IsUpperCase(s string) bool {
	for _, c := range s {
		if !IsUpperRune(c) {
			return false
		}
	}
	return true
}

func LowerCaseFirst(in string) string {
	return strings.ToLower(in[0:1]) + in[1:]
}

func UpperCaseFirst(in string) string {
	return strings.ToUpper(in[0:1]) + in[1:]
}
