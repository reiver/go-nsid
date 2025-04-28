package nsid

// Join combines an NSID domain-authority (such as "com.example") with an NSID name (such as "fooBar") to construct an NSID (such as "com.example.fooBar").
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
//
// Join more-or-less does the opposite of [Split].
//
// Note that the NSID that Join returns is NOT normalized.
// Use [JoinAndNormalize] to create a normalized NSID.
//
// Join does not validate the overall resulting NSID.
// To validate the overall resulting NSID call [Validate].
//
// If you are not sure whether to use Join or [JoinAndNormalize], use [JoinAndNormalize].
func Join(domainAuthority string, name string) string {
	return Construct(domainAuthority, name)
}
