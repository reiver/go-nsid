package nsid

import (
	"fmt"
)

// Join combines an NSID domain-authority (such as "com.example") with an NSID name (such as "fooBar") to construct an NSID (such as "com.example.fooBar").
//
// Join more-or-less does the opposite of [Split].
//
// Join will normalize the domain-authority and will also validate the overall resulting NSID.
// Join will return an error if the resulting NSID fails validation.
func Join(domainAuthority string, name string) (string, error) {
	str := fmt.Sprintf("%s.%s", domainAuthority, name)
	normalized := Normalize(str)
        return normalized, Validate(normalized)
}

