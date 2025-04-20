package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleConstruct() {

	nsID := nsid.Construct("org.archive", "video.streaming", "createStream")

	fmt.Printf("nsid: %s\n", nsID)

	// Output:
	// nsid: org.archive.video.streaming.createStream
}
