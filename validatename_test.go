package nsid

import (
	"testing"

	"math/rand"

	"github.com/reiver/go-arbitrary"
)

func TestValidateName(t *testing.T) {

	tests := []struct{
		Name string
	}{
		{
			Name: "fooBar",
		},
		{
			Name: "ping",
		},
		{
			Name: "c",
		},
		{
			Name: "stuff",
		},
		{
			Name: "thing",
		},



		{
			Name: "a",
		},
		{
			Name: "ab",
		},
		{
			Name: "abc",
		},
		{
			Name: "abcd",
		},
		{
			Name: "abcde",
		},
		{
			Name: "abcdef",
		},
		{
			Name: "abcdefg",
		},
		{
			Name: "abcdefgh",
		},
		{
			Name: "abcdefghi",
		},
		{
			Name: "abcdefghij",
		},
		{
			Name: "abcdefghijk",
		},
		{
			Name: "abcdefghijkl",
		},
		{
			Name: "abcdefghijklm",
		},
		{
			Name: "abcdefghijklmn",
		},
		{
			Name: "abcdefghijklmno",
		},
		{
			Name: "abcdefghijklmnop",
		},
		{
			Name: "abcdefghijklmnopq",
		},
		{
			Name: "abcdefghijklmnopqr",
		},
		{
			Name: "abcdefghijklmnopqrs",
		},
		{
			Name: "abcdefghijklmnopqrst",
		},
		{
			Name: "abcdefghijklmnopqrstu",
		},
		{
			Name: "abcdefghijklmnopqrstuv",
		},
		{
			Name: "abcdefghijklmnopqrstuvw",
		},
		{
			Name: "abcdefghijklmnopqrstuvwx",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxy",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzA",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzAB",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABC",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCD",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDE",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEF",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFG",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGH",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHI",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJ",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJK",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKL",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLM",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMN",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOP",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQ",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQR",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRS",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTU",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVW",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWX",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZab",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabc",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcd",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcde",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdef",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefg",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefgh",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghi",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij",
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijk",
		},
	}

	{
		var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

		for i:=0; i<0xFF; i++ {
			var length int = 1 + rand.Intn(63)

			var name string = arbitrary.String(length, chars)

			test := struct{
				Name string
			}{
				Name: name,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		var doesNotMatter string = "localhost.once.twice.thrice.fource"+"."+test.Name

		err := validateName(test.Name, doesNotMatter)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME: %q", test.Name)
			continue
		}

	}
}

func TestValidateName_fail(t *testing.T) {

	tests := []struct{
		Name string
		DomainAuthority string
		Expected string
	}{
		{
			Name: "",
			DomainAuthority:"com.example",
			Expected: `nsid: nsid name ("") of nsid ("com.example.") is less-than 1 character`,
		},
		{
			Name: "abcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyz",
			DomainAuthority:"com.example",
			Expected: `nsid: nsid name ("abcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyz") of nsid ("com.example.abcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyzAbcdefgHijklmnopQrsTuvWxyz") is greater-than 63 character`,
		},



		{
			Name: "ping14",
			DomainAuthority:"net.users.bob",
			Expected: `nsid: character №4 ('1') (U+0031) of name ("ping14") of nsid ("net.users.bob.ping14") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Name: "c-1",
			DomainAuthority:"a-0.b-1",
			Expected: `nsid: character №1 ('-') (U+002D) of name ("c-1") of nsid ("a-0.b-1.c-1") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Name: "stuff.txt",
			DomainAuthority:"a.b",
			Expected: `nsid: character №5 ('.') (U+002E) of name ("stuff.txt") of nsid ("a.b.stuff.txt") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},
		{
			Name: "thing[2]",
			DomainAuthority:"cn.8.lex",
			Expected: `nsid: character №5 ('[') (U+005B) of name ("thing[2]") of nsid ("cn.8.lex.thing[2]") is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')`,
		},



		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkl",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkl") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkl") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklm",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklm") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklm") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmln",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmln") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmln") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlno",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlno") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlno") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnop",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnop") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnop") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopq",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopq") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopq") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqr",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqr") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqr") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrs",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrs") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrs") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrst",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrst") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrst") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstu",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstu") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstu") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuv",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuv") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuv") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvw",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvw") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvw") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwx",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwx") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwx") is greater-than 63 character`,
		},
		{
			Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwxy",
			DomainAuthority: "c.e",
			Expected: `nsid: nsid name ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwxy") of nsid ("c.e.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmlnopqrstuvwxy") is greater-than 63 character`,
		},
	}

	for testNumber, test := range tests {

		var nsidString string = test.DomainAuthority+"."+test.Name

		err := validateName(test.Name, nsidString)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME:             %q", test.Name)
			t.Logf("DOMAIN-AUTHORITY: %q", test.DomainAuthority)
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
				t.Logf("DOMAIN-AUTHORITY: %q", test.DomainAuthority)
				t.Logf("NSID:             %q", nsidString)
				continue
			}
		}

	}
}
