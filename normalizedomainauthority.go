package nsid

// NormalizeDomainAuthority returns the normalized form of an domain-authority, as defined in:
// https://atproto.com/specs/nsid
//
// The domain-authority (such as "com.example") is part of an NSID (such as "com.example.fooBar").
//
// In simple language, you can think of an NSID domain-authority as an Internet domain-name written in reverse-order.
// For example, if the Internet domain-name was "example.com", then the NSID domain-authority would be "com.example".
//
// An example of a non-normalized NSID domain-authority would be "COM.Example".
// Normalizing that non-normalized NSID domain-authority would result in "com.example".
//
// Note that if you want to normalize a whole NSID rather than just a domain-authority, then instead use [Normalize].
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
