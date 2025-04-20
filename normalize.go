package nsid

import (
	"strings"
)

// Normalize returns the normalized form of an NSID, as defined in:
// https://atproto.com/specs/nsid
//
// Normalize does NOT validate the NSID.
// To validate, call [Validate].
//
// An example of a non-normalized NSID domain-authority would be "COM.Example.fooBar".
// Normalizing that non-normalized NSID domain-authority would result in "com.example.fooBar".
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
