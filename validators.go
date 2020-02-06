package avro2go

import (
	"regexp"
)

// https://avro.apache.org/docs/1.8.1/spec.html#enums
var symbolRegex = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)

func validateByRegex(re *regexp.Regexp, s []byte) bool {
	return re.Match(s)
}

func validateSymbol(s []byte) bool {
	return validateByRegex(symbolRegex, s)
}

var nameRegex = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)

func validateName(s []byte) bool {
	return validateByRegex(nameRegex, s)
}
