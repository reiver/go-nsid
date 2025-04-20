package nsid

import (
	"strings"
)

// Construct creates an NSID (Namespaced Identifier) that is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
//
// Construct is similar to [Join], but is more flexible.
//
// Construct allows the domain-authority to be broken into parts.
func Construct(values ...string) string {
	var nsID string = strings.Join(values, ".")
	return Normalize(nsID)
}
