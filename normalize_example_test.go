package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleNormalize() {

	var nsID string = "COM.Example.fooBar"

	normalized := nsid.Normalize(nsID)

	fmt.Printf("original nsid:   %s\n", nsID)
	fmt.Printf("noramlized nsid: %s\n", normalized)

	// Output:
	// original nsid:   COM.Example.fooBar
	// noramlized nsid: com.example.fooBar
}
