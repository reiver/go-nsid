package nsid_test

import (
	"testing"

	"github.com/reiver/go-nsid"
)

func TestValidate(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value: "",
			Expected: "nsid: empty NSID",
		},



		{
			Value: "com.example.😈",
			Expected: `nsid: character №12 ('😈') (U+1F608) of nsid ("com.example.😈") is not an ASCII character`,
		},
		{
			Value: "com.example.a😈b",
			Expected: `nsid: character №13 ('😈') (U+1F608) of nsid ("com.example.a😈b") is not an ASCII character`,
		},
		{
			Value: "com.😈.abc",
			Expected: `nsid: character №4 ('😈') (U+1F608) of nsid ("com.😈.abc") is not an ASCII character`,
		},
		{
			Value: "😈.example.abc",
			Expected: `nsid: character №0 ('😈') (U+1F608) of nsid ("😈.example.abc") is not an ASCII character`,
		},
		{
			Value: "c😈m.example.abc",
			Expected: `nsid: character №1 ('😈') (U+1F608) of nsid ("c😈m.example.abc") is not an ASCII character`,
		},



		{
			Value: "com.example",
			Expected: `nsid: nsid ("com.example") should have at least 3 segments but actually has 2`,
		},
		{
			Value: "a.b",
			Expected: `nsid: nsid ("a.b") should have at least 3 segments but actually has 2`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba") should have at least 3 segments but actually has 2`,
		},
		{
			Value: "com",
			Expected: `nsid: nsid ("com") should have at least 3 segments but actually has 1`,
		},
		{
			Value: "a",
			Expected: `nsid: nsid ("a") should have at least 3 segments but actually has 1`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz") should have at least 3 segments but actually has 1`,
		},



		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCD",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCD") should have at most 317 characters but actually has 318`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDE",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDE") should have at most 317 characters but actually has 319`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEF",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEF") should have at most 317 characters but actually has 320`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOP",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOP") should have at most 317 characters but actually has 330`,
		},
		{
			Value: "abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Expected: `nsid: nsid ("abcdefghijklmnopqrstuvwxyz.zyxwvutsrqponmlkjihgfedcba.ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ") should have at most 317 characters but actually has 340`,
		},
	}

	for testNumber, test := range tests {

		actualError := nsid.Validate(test.Value)

		actual := actualError.Error()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual error message is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
