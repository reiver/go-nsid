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



		{
			Value: "com.example.0",
			Expected: `nsid: character №0 ('0') (U+0030) of name ("0") of nsid ("com.example.0") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.1",
			Expected: `nsid: character №0 ('1') (U+0031) of name ("1") of nsid ("com.example.1") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.2",
			Expected: `nsid: character №0 ('2') (U+0032) of name ("2") of nsid ("com.example.2") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.3",
			Expected: `nsid: character №0 ('3') (U+0033) of name ("3") of nsid ("com.example.3") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.4",
			Expected: `nsid: character №0 ('4') (U+0034) of name ("4") of nsid ("com.example.4") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.5",
			Expected: `nsid: character №0 ('5') (U+0035) of name ("5") of nsid ("com.example.5") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.6",
			Expected: `nsid: character №0 ('6') (U+0036) of name ("6") of nsid ("com.example.6") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.7",
			Expected: `nsid: character №0 ('7') (U+0037) of name ("7") of nsid ("com.example.7") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.8",
			Expected: `nsid: character №0 ('8') (U+0038) of name ("8") of nsid ("com.example.8") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.9",
			Expected: `nsid: character №0 ('9') (U+0039) of name ("9") of nsid ("com.example.9") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},



		{
			Value: "com.example.0xDEADBEEF",
			Expected: `nsid: character №0 ('0') (U+0030) of name ("0xDEADBEEF") of nsid ("com.example.0xDEADBEEF") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.123",
			Expected: `nsid: character №0 ('1') (U+0031) of name ("123") of nsid ("com.example.123") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.2b",
			Expected: `nsid: character №0 ('2') (U+0032) of name ("2b") of nsid ("com.example.2b") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.3three",
			Expected: `nsid: character №0 ('3') (U+0033) of name ("3three") of nsid ("com.example.3three") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.4our",
			Expected: `nsid: character №0 ('4') (U+0034) of name ("4our") of nsid ("com.example.4our") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.5ive",
			Expected: `nsid: character №0 ('5') (U+0035) of name ("5ive") of nsid ("com.example.5ive") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.6ix",
			Expected: `nsid: character №0 ('6') (U+0036) of name ("6ix") of nsid ("com.example.6ix") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.7even",
			Expected: `nsid: character №0 ('7') (U+0037) of name ("7even") of nsid ("com.example.7even") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.8ight",
			Expected: `nsid: character №0 ('8') (U+0038) of name ("8ight") of nsid ("com.example.8ight") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Value: "com.example.9876543210",
			Expected: `nsid: character №0 ('9') (U+0039) of name ("9876543210") of nsid ("com.example.9876543210") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
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
