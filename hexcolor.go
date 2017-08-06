package main

import (
	"errors"
	"regexp"
	"strings"
)

type HexColor struct {
	code string
}

func NewHexColor(code string) HexColor {
	return HexColor{code: strings.ToLower(code)}
}

func (color *HexColor) Verify() []error {
	// Capacity set as 2 since there is only 2 possible errors
	errs := make([]error, 0, 2)

	if !vlen(color.code) {
		errs = append(errs, errors.New("Invalid length of string is given"))
	}
	if !vchars(color.code) {
		errs = append(errs, errors.New("Invalid character have found"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

// Verify length of color code
// Length of code must be 3, shorten length, or 6
func vlen(code string) bool {
	l := len(code)
	return l == 3 || l == 6
}

// Verify the characters of string
// Uses only hex characters
func vchars(code string) bool {
	var rhex *regexp.Regexp = regexp.MustCompile("^[0-9a-f]{3,6}$")
	return rhex.MatchString(code)
}

// ToDo: func (color *HexColor) ToFull() string
// ToDo: func (color *HexColor) Compress() (string, error)
