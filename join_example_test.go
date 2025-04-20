package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleJoin() {

	var domainAuthority string = "com.example"
	var name string            = "fooBar"

	nsID := nsid.Join(domainAuthority, name)

	fmt.Printf("domain-authority: %s\n", domainAuthority)
	fmt.Printf("name:                         %s\n", name)
	fmt.Printf("nsid:             %s\n", nsID)

	// Output:
	// domain-authority: com.example
	// name:                         fooBar
	// nsid:             com.example.fooBar
}
