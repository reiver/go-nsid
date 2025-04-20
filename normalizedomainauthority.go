package nsid

// NormalizeDomainAuthority returns the normalized form of an domain-authority, as defined in:
// https://atproto.com/specs/nsid
//
// The domain-authority is part of an NSID.
//
// For example, if the NSID was:
//
//	ONCE.TWICE.THRICE.FOURCE
//
// Then the domain-authority would be:
//
//	ONCE.TWICE.THRICE
//
// And then the normalized form of the domain-authority would be:
//
//	once.twice.thrice
//
// So, for example::
//
//	var domainAuthority string = "ONCE.TWICE.THRICE"
//
//	var result string = nsid.NormalizeDomainAuthority(domainAuthority)
//	// result == "once.twice.thrice"
func NormalizeDomainAuthority(value string) string {
	const lenbuffer int = 256
	var buffer [lenbuffer]byte

	var p []byte
	{
		var lenvalue int = len(value)

		if lenvalue <= lenbuffer  {
			p = buffer[:len(value)]
		} else {
			p = make([]byte, lenvalue)
		}
	}

	copy(p, value)

	for i,b := range p {
		switch {
		case 'A' <= b && b <= 'Z':
			p[i] = b - 'A' + 'a'
		}
	}

	return string(p)
}
