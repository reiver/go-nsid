package nsid

import (
	"unicode/utf8"

	"github.com/reiver/go-erorr"
)

// Validate returns an error if the NSID is invalid.
// It returns nil if the NSID is valid.
func Validate(value string) error {
	if "" == value {
		return errEmptyNSID
	}

	// "Overall NSID:" "must contain only ASCII characters"
	for charIndex, char := range value {
		if 127 < char {
			return erorr.Errorf("nsid: character №%d (%q) (%U) of nsid (%q) is not an ASCII character", charIndex, char, char, value)
		}
	}

	domainAuthorityParts, name, numSegments := split(value)

	// "Overall NSID:" "must have at least 3 segments"
	if numSegments < 3 {
		return erorr.Errorf("nsid: nsid (%q) should have at least 3 segments but actually has %d", value, numSegments)
	}

	// "Overall NSID:" "can have a maximum total length of 317 characters"
	{
		const atLeast = 317

		var length int = utf8.RuneCountInString(value)

		if atLeast < length {
			return erorr.Errorf("nsid: nsid (%q) should have at least %d characters but actually has %d", value, atLeast, length)
		}
	}

	{
		err := validateDomainAuthority(domainAuthorityParts, value)
		if nil != err {
			return err
		}
	}

	{
		err := validateName(name, value)
		if nil != err {
			return err
		}
	}

	return nil
}
