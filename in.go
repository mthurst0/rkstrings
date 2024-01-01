package rkstrings

import "strings"

// HasAnyPrefix returns true if s has any of the given prefixes
func HasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if len(prefix) > 0 && strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnySuffix returns true if s has any of the given suffixes
func HasAnySuffix(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if len(suffix) > 0 && strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// Contains returns true if v is in a
func Contains(a []string, v string) bool {
	for _, s := range a {
		if s == v {
			return true
		}
	}
	return false
}

// ContainsFold returns true if v is in a, under "simple Unicode case-folding"
func ContainsFold(a []string, v string) bool {
	for _, s := range a {
		if strings.EqualFold(s, v) {
			return true
		}
	}
	return false
}
