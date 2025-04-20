package nsid

// Join combines an NSID domain-authority (such as "com.example") with an NSID name (such as "fooBar") to construct an NSID (such as "com.example.fooBar").
//
// Join more-or-less does the opposite of [Split].
//
// Join will normalize the domain-authority.
//
// Join does not validate the overall resulting NSID.
// To validate the overall resulting NSID call [Validate].
func Join(domainAuthority string, name string) string {
	return Construct(domainAuthority, name)
}

