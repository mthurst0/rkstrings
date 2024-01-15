package rkstrings

import "strings"

func convertWithFilter(s string, filter func(rune) bool, replacement rune) string {
	var result strings.Builder
	for _, r := range s {
		if filter(r) {
			result.WriteRune(r)
		} else {
			result.WriteRune(replacement)
		}
	}
	return result.String()
}

// IsPermittedInPOSIXFilename returns true if the given rune is permitted in a POSIX filename.
func IsPermittedInPOSIXFilename(c rune) bool {
	return IsASCIIAlphaNumericRune(c) || c == '.' || c == '_' || c == '-'
}

// IsPermittedInPOSIXPath returns true if the given rune is permitted in a POSIX path.
func IsPermittedInPOSIXPath(c rune) bool {
	return IsPermittedInPOSIXFilename(c) || c == '/'
}

// ConvertToSafePOSIXFilename converts the given string to a POSIX filename.
func ConvertToSafePOSIXFilename(s string) string {
	return convertWithFilter(s, IsPermittedInPOSIXFilename, '-')
}

// ConvertToSafePOSIXPath converts the given string to a POSIX path.
func ConvertToSafePOSIXPath(s string) string {
	return convertWithFilter(s, IsPermittedInPOSIXPath, '-')
}
