package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedDomainAuthority string
		ExpectedName string
	}{
		{
			Value:                   "",
			ExpectedDomainAuthority: "",
			ExpectedName:             "",
		},



		{
			Value:                   "apple",
			ExpectedDomainAuthority: "apple",
			ExpectedName:                  "",
		},
		{
			Value:                   "apple.banana",
			ExpectedDomainAuthority: "apple.banana",
			ExpectedName:                         "",
		},
		{
			Value:                   "apple.banana.cherry",
			ExpectedDomainAuthority: "apple.banana",
			ExpectedName:                         "cherry",
		},
		{
			Value:                   "apple.banana.cherry.date",
			ExpectedDomainAuthority: "apple.banana.cherry",
			ExpectedName:                                "date",
		},



		{
			Value:                   "ApPlE",
			ExpectedDomainAuthority: "apple",
			ExpectedName:                  "",
		},
		{
			Value:                   "aPpLe.BaNaNa",
			ExpectedDomainAuthority: "apple.banana",
			ExpectedName:                         "",
		},
		{
			Value:                   "ApPlE.bAnAnA.cHeRrY",
			ExpectedDomainAuthority: "apple.banana",
			ExpectedName:                         "cHeRrY",
		},
		{
			Value:                   "aPpLe.BaNaNa.ChErRy.DaTe",
			ExpectedDomainAuthority: "apple.banana.cherry",
			ExpectedName:                                "DaTe",
		},
	}

	for testNumber, test := range tests {

		actualDomainAuthority, actualName := nsid.Split(test.Value)

		{
			actual := actualDomainAuthority
			expected := test.ExpectedDomainAuthority

			if expected != actual {
				t.Errorf("For test #%d, the actual 'domain-authority' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			actual := actualName
			expected := test.ExpectedName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
