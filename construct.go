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
// Note that the NSID that Construct returns is NOT normalized.
// Use [ConstructAndNormalize] to create a normalized NSID.
//
// If you are not sure whether to use Construct or [ConstructAndNormalize], use [ConstructAndNormalize].
func Construct(values ...string) string {
	return strings.Join(values, ".")
}
