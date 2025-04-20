package nsid

// NSID represents a NSID (Namespaced Identifier) that is part of BlueSky's AT-Protocol.
type NSID string

// ConstructNSID creates an NSID (Namespaced Identifier) that is part of BlueSky's AT-Protocol.
func ConstructNSID(values ...string) (NSID, error) {
	nsIDString := Construct(values...)
	var nsID NSID = NSID(nsIDString)

	return nsID, Validate(nsIDString)
}

// MustConstructNSID is similar to [ConstructNSID] expect it panic()s if thre is an error.
func MustConstructNSID(values ...string) NSID {
	value, err := ConstructNSID(values...)
	if nil != err {
		panic(err)
	}

	return value
}

// DomainAuthority returns the domain-authority of an NSID.
func (receiver NSID) DomainAuthority() string {
	value, _ := receiver.Split()
	return value
}

// Name returns the domain-authority of an NSID.
func (receiver NSID) Name() string {
	_, value := receiver.Split()
	return value
}

// Split returns the domain-authority and name of an NSID.
func (receiver NSID) Split() (domainAuthority string, name string) {
	var str string = string(receiver)
	return Split(str)
}

// Validate returns an error if the NSID is invalid.
// It returns nil if the NSID is valid.
func (receiver NSID) Validate() error {
	var str string = string(receiver)
	return Validate(str)
}
