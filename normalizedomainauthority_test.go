package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestNormalizeDomainAuthority(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value:    "",
			Expected: "",
		},



		{
			Value:    "apple",
			Expected: "apple",
		},
		{
			Value:    "apple.banana",
			Expected: "apple.banana",
		},
		{
			Value:    "apple.banana.cherry",
			Expected: "apple.banana.cherry",
		},
		{
			Value:    "apple.banana.cherry.date",
			Expected: "apple.banana.cherry.date",
		},



		{
			Value:    "ApPlE",
			Expected: "apple",
		},
		{
			Value:    "aPpLe.BaNaNa",
			Expected: "apple.banana",
		},
		{
			Value:    "ApPlE.bAnAnA.cHeRrY",
			Expected: "apple.banana.cherry",
		},
		{
			Value:    "aPpLe.BaNaNa.ChErRy.DaTe",
			Expected: "apple.banana.cherry.date",
		},



		{
			Value:    "ONCE.TWICE.THRICE.FOURCE",
			Expected: "once.twice.thrice.fource",
		},
	}

	for testNumber, test := range tests {

		actual := nsid.NormalizeDomainAuthority(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized-domain-authority is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
