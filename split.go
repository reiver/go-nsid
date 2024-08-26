package nsid

import (
	"strings"
)

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
