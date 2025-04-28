package nsid

// ConstructAndNormalize and similar to [Construct] but it also normalizes the resulting NSID (Namespaced Identifier).
//
// ConstructAndNormalize creates an NSID (Namespaced Identifier).
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
//
// ConstructAndNormalize is similar to [Join], but is more flexible.
//
// ConstructAndNormalize allows the domain-authority to be broken into parts.
//
// If you are not sure whether to use [Construct] or ConstructAndNormalize, use ConstructAndNormalize.
func ConstructAndNormalize(values ...string) string {
	return Normalize(Construct(values...))
}
