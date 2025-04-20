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
			Value: "-.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("-") of domain-authority ("-.example") of nsid ("-.example.abc") cannot begin with hyphen ('-')`,
		},
		{
			Value: "-com.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("-com") of domain-authority ("-com.example") of nsid ("-com.example.abc") cannot begin with hyphen ('-')`,
		},
		{
			Value: "com-.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("com-") of domain-authority ("com-.example") of nsid ("com-.example.abc") cannot end with hyphen ('-')`,
		},
		{
			Value: "com.-.abc",
			Expected: `nsid: nsid domain-authority part №1 ("-") of domain-authority ("com.-") of nsid ("com.-.abc") cannot begin with hyphen ('-')`,
		},
		{
			Value: "com.-example.abc",
			Expected: `nsid: nsid domain-authority part №1 ("-example") of domain-authority ("com.-example") of nsid ("com.-example.abc") cannot begin with hyphen ('-')`,
		},
		{
			Value: "com.example-.abc",
			Expected: `nsid: nsid domain-authority part №1 ("example-") of domain-authority ("com.example-") of nsid ("com.example-.abc") cannot end with hyphen ('-')`,
		},



		{
			Value: "0.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("0") of domain-authority ("0.example") of nsid ("0.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "1.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("1") of domain-authority ("1.example") of nsid ("1.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "2.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("2") of domain-authority ("2.example") of nsid ("2.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "3.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("3") of domain-authority ("3.example") of nsid ("3.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "4.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("4") of domain-authority ("4.example") of nsid ("4.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "5.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("5") of domain-authority ("5.example") of nsid ("5.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "6.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("6") of domain-authority ("6.example") of nsid ("6.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "7.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("7") of domain-authority ("7.example") of nsid ("7.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "8.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("8") of domain-authority ("8.example") of nsid ("8.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "9.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("9") of domain-authority ("9.example") of nsid ("9.example.abc") cannot begin with numerical-digit`,
		},



		{
			Value: "0om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("0om") of domain-authority ("0om.example") of nsid ("0om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "1om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("1om") of domain-authority ("1om.example") of nsid ("1om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "2om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("2om") of domain-authority ("2om.example") of nsid ("2om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "3om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("3om") of domain-authority ("3om.example") of nsid ("3om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "4om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("4om") of domain-authority ("4om.example") of nsid ("4om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "5om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("5om") of domain-authority ("5om.example") of nsid ("5om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "6om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("6om") of domain-authority ("6om.example") of nsid ("6om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "7om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("7om") of domain-authority ("7om.example") of nsid ("7om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "8om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("8om") of domain-authority ("8om.example") of nsid ("8om.example.abc") cannot begin with numerical-digit`,
		},
		{
			Value: "9om.example.abc",
			Expected: `nsid: nsid domain-authority part №0 ("9om") of domain-authority ("9om.example") of nsid ("9om.example.abc") cannot begin with numerical-digit`,
		},



		{
			Value: "com.example.",
			Expected: `nsid: nsid name ("") of nsid ("com.example.") is less-than 1 character`,
		},
		{
			Value: "com.example.abcdefghijklmnopqrstuvwxyz012345ABCDEFGHIJKLMNOPQRSTUVWXYZ543210",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyz012345ABCDEFGHIJKLMNOPQRSTUVWXYZ543210") of nsid ("com.example.abcdefghijklmnopqrstuvwxyz012345ABCDEFGHIJKLMNOPQRSTUVWXYZ543210") is greater-than 63 character`,
		},



		{
			Value: "com.example.🙂",
			Expected: `nsid: character №12 ('🙂') (U+1F642) of nsid ("com.example.🙂") is not an ASCII character`,
		},
		{
			Value: "com.example.a🙂c",
			Expected: `nsid: character №13 ('🙂') (U+1F642) of nsid ("com.example.a🙂c") is not an ASCII character`,
		},



		{
			Value: "com.example.ED7BA470-8E54-465E-825C-99712043E01C",
			Expected: `nsid: character №8 ('-') (U+002D) of name ("ED7BA470-8E54-465E-825C-99712043E01C") of nsid ("com.example.ED7BA470-8E54-465E-825C-99712043E01C") cannot be a hyphen but must instead be an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z'), or digits ('0'-'9')`,
		},



		{
			Value: "com.example.0",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("0") of nsid ("com.example.0") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.1",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("1") of nsid ("com.example.1") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.2",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("2") of nsid ("com.example.2") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.3",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("3") of nsid ("com.example.3") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.4",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("4") of nsid ("com.example.4") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.5",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("5") of nsid ("com.example.5") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.6",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("6") of nsid ("com.example.6") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.7",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("7") of nsid ("com.example.7") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.8",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("8") of nsid ("com.example.8") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.9",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("9") of nsid ("com.example.9") cannot be a digit ('0'-'9')`,
		},



		{
			Value: "com.example.0xDEADBEEF",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("0xDEADBEEF") of nsid ("com.example.0xDEADBEEF") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.123",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("123") of nsid ("com.example.123") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.2b",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("2b") of nsid ("com.example.2b") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.3three",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("3three") of nsid ("com.example.3three") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.4our",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("4our") of nsid ("com.example.4our") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.5ive",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("5ive") of nsid ("com.example.5ive") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.6ix",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("6ix") of nsid ("com.example.6ix") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.7even",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("7even") of nsid ("com.example.7even") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.8ight",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("8ight") of nsid ("com.example.8ight") cannot be a digit ('0'-'9')`,
		},
		{
			Value: "com.example.9876543210",
			Expected: `nsid: the first character (i.e., character №0) of nsid name ("9876543210") of nsid ("com.example.9876543210") cannot be a digit ('0'-'9')`,
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
