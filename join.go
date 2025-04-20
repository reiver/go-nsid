package nsid

import (
	"fmt"
)

// Join combined an NSID domain-authority and an NSID name to construct an NSID.
//
// Join more-or-less does the opposite of [Split].
//
// Join will normalize the domain-authority and also validate the overall NSID.
// Join will return an error if the resulting NSID fails validation.
func Join(domainAuthority string, name string) (string, error) {
	str := fmt.Sprintf("%s.%s", domainAuthority, name)
        return Normalize(str), Validate(str)
}

