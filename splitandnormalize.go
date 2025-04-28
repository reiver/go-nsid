package nsid

// Split returns the domain-authority and name of an NSID, as defined here:
// https://atproto.com/specs/nsid
//
// Split normalizes the domain-authority of the NSID that it returns.
//
// If you are not sure whether to use [Split] or SpliAndNormalize, use SplitAndNormalize.
func SplitAndNormalize(value string) (domainAuthority string, name string) {
	domainAuthority, name = Split(value)

	domainAuthority = NormalizeDomainAuthority(domainAuthority)

	return domainAuthority, name
}
