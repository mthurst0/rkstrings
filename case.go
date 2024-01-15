package rkstrings

import (
	"strings"
	"unicode"
)

// ToCamelCaseLowercaseFirst converts a string to camelCase.
func ToCamelCaseLowercaseFirst(in string) string {
	return LowerCaseFirst(ToCamelCase(in))
}

// ToCamelCase converts a string to CamelCase.
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

// backUpOneBuilder wraps a strings.Builder that can back up one rune.
type backUpOneBuilder struct {
	sb       strings.Builder
	lastRune rune
}

func (b *backUpOneBuilder) WriteRune(r rune) {
	if b.lastRune != 0 {
		b.sb.WriteRune(b.lastRune)
	}
	b.lastRune = r
}

func (b *backUpOneBuilder) BackItUp() {
	b.lastRune = 0
}

func (b *backUpOneBuilder) String() string {
	b.sb.WriteRune(b.lastRune)
	return b.sb.String()
}

// convertFromCamelCase converts a camelCase string to a delimited string.
// The delimiter is specified by the delim parameter.
func convertFromCamelCase(s string, delim rune) string {
	var newName backUpOneBuilder
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

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(in string) string {
	return convertFromCamelCase(ToCamelCase(in), '_')
}

// ToDashCase converts a string to dash-case.
func ToDashCase(in string) string {
	return convertFromCamelCase(ToCamelCase(in), '-')
}

// IsLowerRune returns true if the rune is a letter and is lower case.
func IsLowerRune(r rune) bool {
	return unicode.IsLower(r) && unicode.IsLetter(r)
}

// IsUpperRune returns true if the rune is a letter and is upper case.
func IsUpperRune(r rune) bool {
	return unicode.IsUpper(r) && unicode.IsLetter(r)
}

// IsLowerCase returns true if the string is all lower case letters.
func IsLowerCase(s string) bool {
	for _, c := range s {
		if !IsLowerRune(c) {
			return false
		}
	}
	return true
}

// IsUpperCase returns true if the string is all upper case letters.
func IsUpperCase(s string) bool {
	for _, c := range s {
		if !IsUpperRune(c) {
			return false
		}
	}
	return true
}

// LowerCaseFirst returns the string with the first letter lower cased.
func LowerCaseFirst(in string) string {
	return strings.ToLower(in[0:1]) + in[1:]
}

// UpperCaseFirst returns the string with the first letter upper cased.
func UpperCaseFirst(in string) string {
	return strings.ToUpper(in[0:1]) + in[1:]
}
