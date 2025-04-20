package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleValidate() {

	var nsID string = "com.example.foo-bar"

	err := nsid.Validate(nsID)

	fmt.Printf("nsid: %s\n", nsID)
	fmt.Printf("validation error: %s\n", err)

	// Output:
	// nsid: com.example.foo-bar
	// validation error: nsid: character №3 ('-') (U+002D) of name ("foo-bar") of nsid ("com.example.foo-bar") cannot be a hyphen but must instead be an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z'), or digits ('0'-'9')
}
