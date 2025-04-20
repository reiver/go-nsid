package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestJoin(t *testing.T) {

	tests := []struct{
		DomainAuthority string
		Name string
		Expected string
	}{
		{
			DomainAuthority: "",
			Name: "",
			Expected: ".",
		},



		{
			DomainAuthority: "com.example",
			Name: "fooBar",
			Expected: "com.example.fooBar",
		},
		{
			DomainAuthority: "COM.Example",
			Name: "fooBar",
			Expected: "com.example.fooBar",
		},



		{
			DomainAuthority: "org.archive.video",
			Name: "clipVideo",
			Expected: "org.archive.video.clipVideo",
		},
		{
			DomainAuthority: "Org.Archive.VIDEO",
			Name: "clipVideo",
			Expected: "org.archive.video.clipVideo",
		},
	}

	for testNumber, test := range tests {

		actual := nsid.Join(test.DomainAuthority, test.Name)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual resulting NSID is not what was expected." , testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("DOMAIN-AUTHORITY: %q", test.DomainAuthority)
			t.Logf("NAME: %q", test.Name)
			continue
		}
	}
}
