package nsid

import (
	"strings"
)

// Normalize return the normalized for of an NSID.
//
// Normalize does NOT validate the NSID.
//
// To validate, call [Validate].
func Normalize(value string) string {
	var val NSID = NSID(value)

	domainAuthority, name := val.Split()

	domainAuthority = strings.ToLower(domainAuthority)

	var str string = strings.Join([]string{domainAuthority, name}, ".")

	return str
}
