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

	// "Name:" "the first character can not be a digit"
	{
		var char byte = name[0]

		switch {
		case '0' <= char && char <= '9':
			return erorr.Errorf("nsid: the first character (i.e., character №0) of nsid name (%q) of nsid (%q) cannot be a digit ('0'-'9')", name, nsidString)
		}
	}

	// "Name:" "the allowed characters are ASCII letters only (A-Z, a-z)"
	//
	// "Name:" "digits and hyphens are not allowed"
	{
		for charIndex, char := range name {
			switch {
			case '0' <= char && char <= '9':
				// nothing here
			case 'A' <= char && char <= 'Z':
				// nothing here
			case 'a' <= char && char <= 'z':
				// nothing here
			default:
				switch {
				case '-' == char:
					return erorr.Errorf("nsid: character №%d (%q) (%U) of name (%q) of nsid (%q) cannot be a hyphen but must instead be an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z'), or digits ('0'-'9')", charIndex, char, char, name, nsidString)
				default:
					return erorr.Errorf("nsid: character №%d (%q) (%U) of name (%q) of nsid (%q) is not an upper-case letter ('A'-'Z'), or lower-case letter ('a'-'z'), or digits ('0'-'9')", charIndex, char, char, name, nsidString)
				}
			}
		}
	}

	return nil
}

