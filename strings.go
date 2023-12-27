package rkstrings

import (
	"strings"
	"unicode"
)

func IsLowerCaseString(s string) bool {
	for _, c := range s {
		if !IsAlphaNumeric(c) {
			return false
		}
	}
	return true
}

func IsAlphaNumeric(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func IsPermittedInPOSIXFilename(c rune) bool {
	return IsAlphaNumeric(c) || c == '.' || c == '_' || c == '-'
}

func IsPermittedInPOSIXPath(c rune) bool {
	return IsPermittedInPOSIXFilename(c) || c == '/'
}

func ConvertToSafePOSIXFilename(src string) string {
	return convertWithFilter(src, IsPermittedInPOSIXFilename, '-')
}

func ConvertToSafePOSIXPath(src string) string {
	return convertWithFilter(src, IsPermittedInPOSIXPath, '-')
}

func LowerCaseFirst(in string) string {
	return strings.ToLower(in[0:1]) + in[1:]
}

func UpperCaseFirst(in string) string {
	return strings.ToUpper(in[0:1]) + in[1:]
}

func IsUpperRune(r rune) bool {
	return unicode.IsUpper(r) && unicode.IsLetter(r)
}

func IsLowerRune(r rune) bool {
	return unicode.IsLower(r) && unicode.IsLetter(r)
}

func ToDashCase(in string) string {
	return ConvertCamelCase(in, '-')
}

func ToSnakeCase(in string) string {
	return ConvertCamelCase(in, '_')
}

// Contains returns true if v is in a
func Contains(a []string, v string) bool {
	for _, value := range a {
		if value == v {
			return true
		}
	}
	return false
}

// HasAnySuffix returns true if s has any of the given suffixes
func HasAnySuffix(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func BoolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func InSlice(list []string, value string) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}
	return false
}

func InSliceCaseInsensitive(list []string, value string) bool {
	for _, v := range list {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}
