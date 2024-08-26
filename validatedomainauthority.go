package nsid

import (
	"strings"
	"unicode/utf8"

	"github.com/reiver/go-erorr"
)

func validateDomainAuthority(domainAuthorityParts []string, nsidString string) error {

	domainAuthorityString := strings.Join(domainAuthorityParts, ".")

	// "Domain authority:" "at most 253 characters (including periods)"
	{
		var length int = utf8.RuneCountInString(domainAuthorityString)

		if 253 < length {
			return erorr.Errorf("nsid: nsid domain-authority (%q) of nsid (%q) character count is %d characters but should not be greater-than 253 characters", domainAuthorityString, nsidString, length)
		}
	}

	// "Domain authority:" "must contain at least two segments"
	{
		var length int = len(domainAuthorityParts)

		if length < 2 {
			return erorr.Errorf("nsid: nsid domain-authority (%q) of nsid (%q) expected to have 2 or more segments but only has %d", domainAuthorityString, nsidString, length)
		}
	}

	// "Domain authority:" "each segment must have at least 1 and at most 63 characters (not including any periods)"
	//
	// "Domain authority:" "the allowed characters are ASCII letters (a-z), digits (0-9), and hyphens (-)"
	//
	// "Domain authority:" "segments can not start or end with a hyphen"
	//
	// "Domain authority:" "the first segment (the top-level domain) can not start with a numeric digit"
	for i, part := range domainAuthorityParts {
		var length int = utf8.RuneCountInString(part)

		if length < 1 {
			return erorr.Errorf("nsid: nsid domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) is less-than 1 character", i, part, domainAuthorityString, nsidString)
		}
		if 63 < length {
			return erorr.Errorf("nsid: nsid domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) is greater-than 63 character", i, part, domainAuthorityString, nsidString)
		}

		for charIndex, char := range part {
			switch {
			case '0' <= char && char <= '9':
				// nothing here
			case 'a' <= char && char <= 'z':
				// nothing here
			case '-' == char:
				// nothing here.
			default:
				return erorr.Errorf("nsid: character №%d (%q) (%U) of domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')", charIndex, char, char, i, part, domainAuthorityString, nsidString)
			}
		}

		if '-' == part[0] {
			return erorr.Errorf("nsid: nsid domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) cannot begin with hyphen ('-')", i, part, domainAuthorityString, nsidString)
		}

		if '-' == part[len(part)-1] {
			return erorr.Errorf("nsid: nsid domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) cannot end with hyphen ('-')", i, part, domainAuthorityString, nsidString)
		}

		if 0 == i {
			var char0 byte = part[0]
			if '0' <= char0 && char0 <= '9' {
				return erorr.Errorf("nsid: nsid domain-authority part №%d (%q) of domain-authority (%q) of nsid (%q) cannot begin with numerical-digit", i, part, domainAuthorityString, nsidString)
			}
		}
	}

	return nil
}

