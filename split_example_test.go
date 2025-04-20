package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleSplit() {

	var nsID string = "COM.Example.fooBar"

	domainAuthority, name := nsid.Split(nsID)

	fmt.Printf("nsid:             %s\n", nsID)
	fmt.Printf("domain-authority: %s\n", domainAuthority)
	fmt.Printf("name:                         %s\n", name)

	// Output:
	// nsid:             COM.Example.fooBar
	// domain-authority: com.example
	// name:                         fooBar
}
