package nsid

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
