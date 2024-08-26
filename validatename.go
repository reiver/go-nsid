package nsid

import (
	"github.com/reiver/go-erorr"
)

func validateName(name string, nsidString string) error {

	// "Name:" "must have at least 1 and at most 63 characters"
	{
		var length int = len(name)

		if length < 1 {
			return erorr.Errorf("nsid: nsid name (%q) of nsid (%q) is less-than 1 character", name, nsidString)
		}
		if 63 < length {
			return erorr.Errorf("nsid: nsid name (%q) of nsid (%q) is greater-than 63 character", name, nsidString)
		}
	}

	// "Name:" "the allowed characters are ASCII letters only (A-Z, a-z)"
	//
	// "Name:" "digits and hyphens are not allowed"
	{
		for charIndex, char := range name {
			switch {
			case 'A' <= char && char <= 'Z':
				// nothing here
			case 'a' <= char && char <= 'z':
				// nothing here
			default:
				return erorr.Errorf("nsid: character №%d (%q) (%U) of name (%q) of nsid (%q) is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z')", charIndex, char, char, name, nsidString)
			}
		}
	}

	return nil
}

