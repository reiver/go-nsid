package nsid

import (
	"fmt"
)

// Join combined an NSID domain-authority and an NSID name to construct an NSID.
//
// Join more-or-less does the opposite of [Split].
func Join(domainAuthority string, name string) (string, error) {
	str := fmt.Sprintf("%s.%s", domainAuthority, name)
        return Normalize(str), Validate(str)
}

