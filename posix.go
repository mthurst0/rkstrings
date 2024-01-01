package rkstrings

import "strings"

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

func IsPermittedInPOSIXFilename(c rune) bool {
	return IsASCIIAlphaNumericRune(c) || c == '.' || c == '_' || c == '-'
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
