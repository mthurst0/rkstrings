package rkstrings

import "regexp"

var uuidRegex *regexp.Regexp

func init() {
	uuidRegex = regexp.MustCompile(
		`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)
}

func IsUUID(s string) bool {
	return uuidRegex.Match([]byte(s))
}
