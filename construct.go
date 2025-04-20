package nsid

import (
	"strings"
)

// Construct creates an NSID (Namespaced Identifier).
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
//
// Construct is similar to [Join], but is more flexible.
//
// Construct allows the domain-authority to be broken into parts.
//
// Note that the NSID that Construct returns is normalized.
func Construct(values ...string) string {
	var nsID string = strings.Join(values, ".")
	return Normalize(nsID)
}
