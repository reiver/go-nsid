package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleNormalizeDomainAuthority() {

	var domainAuthority string = "COM.Example"

	normalized := nsid.NormalizeDomainAuthority(domainAuthority)

	fmt.Printf("original domain-authority:   %s\n", domainAuthority)
	fmt.Printf("noramlized domain-authority: %s\n", normalized)

	// Output:
	// original domain-authority:   COM.Example
	// noramlized domain-authority: com.example
}
