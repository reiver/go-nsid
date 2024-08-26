package nsid

import (
	"testing"

	"math/rand"
	"strings"
	"slices"

	"github.com/reiver/go-arbitrary"
)

func TestValidateDomainAuthority(t *testing.T) {

	tests := []struct{
		Name string
		DomainAuthority []string
	}{
		{
			DomainAuthority: []string{"com","example"},
			Name: "fooBar",
		},
		{
			DomainAuthority: []string{"net","users","bob"},
			Name: "ping",
		},
		{
			DomainAuthority: []string{"a-0","b-1"},
			Name: "c",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "c",
		},
		{
			DomainAuthority: []string{"cn","8","lex"},
			Name: "c",
		},



		{
			DomainAuthority: []string{"a","b"},
			Name: "a",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "ab",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abc",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcd",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcde",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdef",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefg",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefgh",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghi",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghij",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijk",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijkl",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklm",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmn",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmno",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnop",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopq",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqr",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrs",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrst",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstu",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuv",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvw",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwx",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxy",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzA",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzAB",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABC",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCD",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDE",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEF",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFG",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGH",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHI",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJ",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJK",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKL",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLM",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOP",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQ",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQR",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRS",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTU",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVW",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWX",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZab",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabc",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcd",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcde",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdef",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefg",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefgh",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij",
		},
		{
			DomainAuthority: []string{"a","b"},
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijk",
		},
	}

	{
		var namechars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

		for i:=0; i<0xFF; i++ {
			var namelength int = 1 + rand.Intn(63)
			var name string = arbitrary.String(namelength, namechars)

			var domainAuthority []string = strings.Split(arbitrary.InternetHostName(), ".")
			slices.Reverse(domainAuthority)

			test := struct{
				Name string
				DomainAuthority []string
			}{
				Name: name,
				DomainAuthority: domainAuthority,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		var doesNotMatter string = strings.Join(test.DomainAuthority,".")+"."+test.Name

		err := validateDomainAuthority(test.DomainAuthority, doesNotMatter)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME:             %q", test.Name)
			t.Logf("DOMAIN-AUTHORITY: %#v", test.DomainAuthority)
			t.Logf("NSID:             %q", doesNotMatter)
			continue
		}
	}
}

func TestValidateDomainAuthority_fail(t *testing.T) {

	tests := []struct{
		Name string
		DomainAuthority []string
		Expected string
	}{
		{
			Name: "n",
			DomainAuthority:[]string{
				"da34567890"+
				"1234567890"+
				"1234567890"+
				"1234567890"+
				"123456789" ,

				"1234567890"+
				"1234567890"+
				"1234567890"+
				"1234567890"+
				"123456789" ,

				"1234567890"+
				"1234567890"+
				"1234567890"+
				"1234567890"+
				"123456789" ,

				"1234567890"+
				"1234567890"+
				"1234567890"+
				"1234567890"+
				"123456789" ,

				"1234567890"+
				"1234567890"+
				"1234567890"+
				"1234567890"+
				"123456789" ,

				"1234",
			},
			Expected: `nsid: nsid domain-authority ("da34567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234") of nsid ("da34567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234567890123456789012345678901234567890123456789.1234.n") character count is 254 characters but should not be greater-than 253 characters`,
		},



		{
			Name: "n",
			DomainAuthority:[]string{},
			Expected: `nsid: nsid domain-authority ("") of nsid (".n") expected to have 2 or more segments but only has 0`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{""},
			Expected: `nsid: nsid domain-authority ("") of nsid (".n") expected to have 2 or more segments but only has 1`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"once"},
			Expected: `nsid: nsid domain-authority ("once") of nsid ("once.n") expected to have 2 or more segments but only has 1`,
		},



		{
			Name: "n",
			DomainAuthority:[]string{""},
			Expected: `nsid: nsid domain-authority ("") of nsid (".n") expected to have 2 or more segments but only has 1`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"",""},
			Expected: `nsid: nsid domain-authority part №0 ("") of domain-authority (".") of nsid ("..n") is less-than 1 character`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"a",""},
			Expected: `nsid: nsid domain-authority part №1 ("") of domain-authority ("a.") of nsid ("a..n") is less-than 1 character`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"","b"},
			Expected: `nsid: nsid domain-authority part №0 ("") of domain-authority (".b") of nsid (".b.n") is less-than 1 character`,
		},



		{
			Name: "n",
			DomainAuthority:[]string{"A","b"},
			Expected: `nsid: character №0 ('A') (U+0041) of domain-authority part №0 ("A") of domain-authority ("A.b") of nsid ("A.b.n") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"a","B"},
			Expected: `nsid: character №0 ('B') (U+0042) of domain-authority part №1 ("B") of domain-authority ("a.B") of nsid ("a.B.n") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"a","b","C"},
			Expected: `nsid: character №0 ('C') (U+0043) of domain-authority part №2 ("C") of domain-authority ("a.b.C") of nsid ("a.b.C.n") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"a","b$c"},
			Expected: `nsid: character №1 ('$') (U+0024) of domain-authority part №1 ("b$c") of domain-authority ("a.b$c") of nsid ("a.b$c.n") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')`,
		},



		{
			Name: "n",
			DomainAuthority:[]string{"a","b-c", "-def", "g"},
			Expected: `nsid: nsid domain-authority part №2 ("-def") of domain-authority ("a.b-c.-def.g") of nsid ("a.b-c.-def.g.n") cannot begin with hyphen ('-')`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"a","b-c", "def-", "g"},
			Expected: `nsid: nsid domain-authority part №2 ("def-") of domain-authority ("a.b-c.def-.g") of nsid ("a.b-c.def-.g.n") cannot end with hyphen ('-')`,
		},



		{
			Name: "n",
			DomainAuthority:[]string{"0a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("0a") of domain-authority ("0a.bc") of nsid ("0a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"1a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("1a") of domain-authority ("1a.bc") of nsid ("1a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"2a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("2a") of domain-authority ("2a.bc") of nsid ("2a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"3a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("3a") of domain-authority ("3a.bc") of nsid ("3a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"4a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("4a") of domain-authority ("4a.bc") of nsid ("4a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"5a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("5a") of domain-authority ("5a.bc") of nsid ("5a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"6a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("6a") of domain-authority ("6a.bc") of nsid ("6a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"7a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("7a") of domain-authority ("7a.bc") of nsid ("7a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"8a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("8a") of domain-authority ("8a.bc") of nsid ("8a.bc.n") cannot begin with numerical-digit`,
		},
		{
			Name: "n",
			DomainAuthority:[]string{"9a","bc"},
			Expected: `nsid: nsid domain-authority part №0 ("9a") of domain-authority ("9a.bc") of nsid ("9a.bc.n") cannot begin with numerical-digit`,
		},
	}

	for testNumber, test := range tests {

		var nsidString string = strings.Join(test.DomainAuthority, ".")+"."+test.Name

		err := validateDomainAuthority(test.DomainAuthority, nsidString)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME:             %q", test.Name)
			t.Logf("DOMAIN-AUTHORITY: %#v", test.DomainAuthority)
			t.Logf("NSID:             %q", nsidString)
			continue
		}

		{
			expected := test.Expected
			actual := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual error-message is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME:             %q", test.Name)
				t.Logf("DOMAIN-AUTHORITY: %#v", test.DomainAuthority)
				t.Logf("NSID:             %q", nsidString)
				continue
			}
		}

	}
}
