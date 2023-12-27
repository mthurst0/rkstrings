package rkstrings

import (
	"strings"
	"unicode"
)

func convertWithFilter(src string, filter func(rune) bool, replacement rune) string {
	var s strings.Builder
	for _, r := range src {
		if filter(r) {
			s.WriteRune(r)
		} else {
			s.WriteRune(replacement)
		}
	}
	return s.String()
}

func ConvertToSlug(src string) string {
	return strings.ToLower(convertWithFilter(src, IsAlphaNumeric, '-'))
}

type oneHyphenBuilder struct {
	wrapped     strings.Builder
	writeHyphen bool
}

func (b *oneHyphenBuilder) WriteRune(r rune) {
	if r == '-' {
		if b.writeHyphen {
			b.wrapped.WriteRune(r)
		}
		b.writeHyphen = false
	} else {
		b.writeHyphen = true
		b.wrapped.WriteRune(r)
	}
}

func (b *oneHyphenBuilder) String() string {
	return b.wrapped.String()
}

func slugIt(preserveRune func(r rune) bool, s string) string {
	var b oneHyphenBuilder
	firstRune := true
	upperCase := false
	for _, r := range s {
		if IsUpperRune(r) {
			if !firstRune && !upperCase {
				b.WriteRune('-')
			}
			b.WriteRune(unicode.ToLower(r))
			firstRune = false
			upperCase = true
		} else if IsLowerRune(r) {
			b.WriteRune(r)
			firstRune = false
			upperCase = false
		} else if preserveRune(r) {
			b.WriteRune(r)
			firstRune = true
			upperCase = false
		} else {
			if !firstRune {
				b.WriteRune('-')
			}
			firstRune = false
			upperCase = false
		}
	}
	return b.String()
}

// PathSlug translates s into a path string in slug form.
// Where a slug string is all lowercase, with special characters
// replaced by hyphens ('-').
func PathSlug(s string) string {
	return slugIt(func(r rune) bool {
		return r == '/' || r == '\\' || r == '.' || unicode.IsNumber(r)
	}, s)
}

// FullSlug removes all special characters other than '-'
func FullSlug(s string) string {
	return slugIt(func(r rune) bool {
		return unicode.IsNumber(r)
	}, s)
}
