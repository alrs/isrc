// International Standard Recording Code
package isrc

import (
	"fmt"
	"regexp"
	"strconv"
)

const isrcExpStr = "^[a-zA-Z]{2}[a-zA-Z0-9]{3}[0-9]{7}$"

var isrcRegexp *regexp.Regexp

// ISRC represents an International Standard Recording Code
// https://www.dittomusic.com/blog/what-is-an-isrc-code
type ISRC struct {
	country     []rune
	registrant  []rune
	year        uint8
	designation uint32
}

func init() {
	var err error
	// This can only fail if the hard-coded isrcExpStr constant has been
	// modified.
	isrcRegexp, err = regexp.Compile(isrcExpStr)
	if err != nil {
		panic(err)
	}
}

// NewISRC takes a string representation of an ISRC code and returns a pointer
// to an isrc.ISRC and an error variable.
func NewISRC(isrcStr string) (*ISRC, error) {
	var i ISRC
	if !isrcRegexp.Match([]byte(isrcStr)) {
		return &i, fmt.Errorf("invalid ISRC string")
	}
	i.country = []rune(isrcStr[0:2])
	i.registrant = []rune(isrcStr[2:5])

	year, err := strconv.Atoi(isrcStr[5:7])
	if err != nil {
		return &i, fmt.Errorf("error converting release year to integer: %s", err)
	}
	i.year = uint8(year)

	designation, err := strconv.Atoi(isrcStr[7:12])
	if err != nil {
		return &i, fmt.Errorf("error converting designation code to integer: %s",
			err)
	}
	i.designation = uint32(designation)
	return &i, nil
}

// Country returns the two-letter alphabetic country code as a string.
func (i *ISRC) Country() string {
	return string(i.country)
}

// Registrant returns the three-character alphanumeric registrant code
// as a string.
func (i *ISRC) Registrant() string {
	return string(i.registrant)
}

// Year returns the two-digit year code as an unsigned 8-bit integer.
func (i *ISRC) Year() uint8 {
	return i.year
}

// Designation returns the five-digit designation code as an unsigned
// 32-bit integer.
func (i *ISRC) Designation() uint32 {
	return i.designation
}

// String returns an isrc.ISRC as a string.
func (i *ISRC) String() string {
	return fmt.Sprintf("%2s%3s%02d%05d", i.Country(), i.Registrant(), i.year, i.designation)
}
