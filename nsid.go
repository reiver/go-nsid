package nsid

// NSID represents a NSID (Namespaced Identifier).
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
type NSID string

// CreateNSID creates an NSID (Namespaced Identifier).
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func CreateNSID(values ...string) (NSID, error) {
	nsIDString := ConstructAndNormalize(values...)
	var nsID NSID = NSID(nsIDString)

	return nsID, Validate(nsIDString)
}

// MustCreateNSID is similar to [CreateNSID] expect it panic()s if thre is an error.
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func MustCreateNSID(values ...string) NSID {
	value, err := CreateNSID(values...)
	if nil != err {
		panic(err)
	}

	return value
}

// DomainAuthority returns the domain-authority of an NSID.
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func (receiver NSID) DomainAuthority() string {
	value, _ := receiver.Split()
	return value
}

// Name returns the domain-authority of an NSID.
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func (receiver NSID) Name() string {
	_, value := receiver.Split()
	return value
}

// Split returns the domain-authority and name of an NSID.
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func (receiver NSID) Split() (domainAuthority string, name string) {
	return SplitAndNormalize(string(receiver))
}

// Validate returns an error if the NSID is invalid.
// It returns nil if the NSID is valid.
//
// The NSID (Namespaced Identifier) is part of BlueSky's AT-Protocol, as defined here:
// https://atproto.com/specs/nsid
func (receiver NSID) Validate() error {
	return Validate(string(receiver))
}
