package nsid

import (
	"strings"
)

// Normalize returns the normalized form of an NSID.
//
// Normalize does NOT validate the NSID.
// To validate, call [Validate].
func Normalize(value string) string {

	var name string
	var domainAuthorityParts []string
	domainAuthorityParts, name, _ = split(value)
	var domainAuthority string = strings.Join(domainAuthorityParts, ".")

	domainAuthority = NormalizeDomainAuthority(domainAuthority)

	var str string = domainAuthority
	if "" != name {
		str = strings.Join([]string{domainAuthority, name}, ".")
	}

	return str
}
