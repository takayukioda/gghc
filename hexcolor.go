package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type HexColor struct {
	code string
}

func NewHexColor(code string) (*HexColor, []error) {
	color := &HexColor{code: strings.ToLower(code)}
	if errs := verify(color); errs != nil {
		return nil, errs
	}
	return color, nil
}

func (color *HexColor) GetCode() string {
	return color.code
}

/**
 * Returns full length color code
 */
func (color *HexColor) ToFull() string {
	if len(color.code) == 6 {
		return color.code
	}
	rgb := strings.Split(color.code, "")
	return fmt.Sprintf("%s%s%s%s%s%s", rgb[0], rgb[0], rgb[1], rgb[1], rgb[2], rgb[2])
}

/**
 * Compress returns compressed color code
 * color code must be in form of `rrggbb` which each r, g, b is repeated
 * e.g. ff11aa
 * fe11aa will return error with empty text
 */
func (color *HexColor) Compress() (string, error) {
	if len(color.code) == 3 {
		return color.code, nil
	}
	rrggbb := strings.Split(color.code, "")

	if rrggbb[0] != rrggbb[1] || rrggbb[2] != rrggbb[3] || rrggbb[4] != rrggbb[5] {
		return "", errors.New("To compress, rgb have to have same value for each rgb")
	}
	return fmt.Sprintf("%s%s%s", rrggbb[0], rrggbb[2], rrggbb[4]), nil
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

func verify(color *HexColor) []error {
	// Initialize slice with capacity of 2; there is only 2 possible errors
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
