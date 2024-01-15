package rkstrings

// IsASCIIAlphaRune returns true if the rune is an ASCII letter.
func IsASCIIAlphaRune(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

// IsASCIIAlpha returns true if the string is ASCII letters.
func IsASCIIAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsASCIIAlphaRune(c) {
			return false
		}
	}
	return true
}

// IsASCIIAlphaNumericRune returns true if the rune is an ASCII letter or number.
func IsASCIIAlphaNumericRune(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}

// IsASCIIAlphaNumeric returns true if the string is ASCII letters or numbers.
func IsASCIIAlphaNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsASCIIAlphaNumericRune(c) {
			return false
		}
	}
	return true
}

// IsASCIINumericRune returns true if the rune is an ASCII number.
func IsASCIINumericRune(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsASCIINumeric returns true if the string is ASCII numbers.
func IsASCIINumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsASCIINumericRune(c) {
			return false
		}
	}
	return true
}
