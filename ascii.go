package rkstrings

func IsASCIIAlphaRune(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func IsASCIIAlpha(s string) bool {
	for _, c := range s {
		if !IsASCIIAlphaRune(c) {
			return false
		}
	}
	return true
}

func IsASCIIAlphaNumericRune(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func IsASCIIAlphaNumeric(s string) bool {
	for _, c := range s {
		if !IsASCIIAlphaNumericRune(c) {
			return false
		}
	}
	return true
}
