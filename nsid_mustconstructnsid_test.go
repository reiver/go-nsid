package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestMustConstructNSID(t *testing.T) {
	tests := []struct{
		Parts []string
		Expected string
	}{
		{
			Parts: []string{"com","example", "name"},
			Expected:       "com.example.name",
		},
		{
			Parts: []string{"com.example", "name"},
			Expected:       "com.example.name",
		},
		{
			Parts: []string{"com.example.name"},
			Expected:       "com.example.name",
		},



		{
			Parts: []string{"once.twice.thrice.fource"},
			Expected:       "once.twice.thrice.fource",
		},
	}

	for testNumber, test := range tests {
		actual := nsid.MustConstructNSID(test.Parts...)

		expectedString := test.Expected
		expected := nsid.NSID(expectedString)

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PARTS: %#v", test.Parts)
			continue
		}
	}
}
