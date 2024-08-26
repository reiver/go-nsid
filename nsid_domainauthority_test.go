package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestNSID_DomainAuthority(t *testing.T) {
	tests := []struct{
		NSID nsid.NSID
		Expected string
	}{
		{
			NSID: nsid.MustConstructNSID([]string{"com","example", "name"}...),
			Expected:                             "com.example",
		},
		{
			NSID: nsid.MustConstructNSID([]string{"com.example", "name"}...),
			Expected:                             "com.example",
		},
		{
			NSID: nsid.MustConstructNSID([]string{"com.example.name"}...),
			Expected:                             "com.example",
		},



		{
			NSID: nsid.MustConstructNSID([]string{"once.twice.thrice.fource"}...),
			Expected:                             "once.twice.thrice",
		},
	}

	for testNumber, test := range tests {
		actual := test.NSID.DomainAuthority()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("NSID: %#v", test.NSID)
			continue
		}
	}
}
