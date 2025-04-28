package nsid_test

import (
	"fmt"

	"github.com/reiver/go-nsid"
)

func ExampleConstructAndNormalize() {

	nsID := nsid.ConstructAndNormalize("org.archive", "video.streaming", "createStream")

	fmt.Printf("nsid: %s\n", nsID)

	// Output:
	// nsid: org.archive.video.streaming.createStream
}
