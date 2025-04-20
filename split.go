package nsid

import (
	"strings"
)

// Split returns the domain-authority and name of an NSID, as defined here:
// https://atproto.com/specs/nsid
//
// Split normalizes the domain-authority of the NSID that it returns.
func Split(value string) (domainAuthority string, name string) {
	var domainAuthorityParts []string
	domainAuthorityParts, name, _ = split(value)

	domainAuthority = strings.Join(domainAuthorityParts, ".")
	domainAuthority = NormalizeDomainAuthority(domainAuthority)

	return domainAuthority, name
}

func split(value string) (domainAuthority []string, name string, numSegments int) {
	var parts []string = strings.Split(value, ".")

	domainAuthority = parts

	numSegments = len(parts)

	if 3 <= numSegments {
		var length int = len(parts)

		name = parts[length-1]
		domainAuthority = parts[:length-1]
	}

	return domainAuthority, name, numSegments
}
