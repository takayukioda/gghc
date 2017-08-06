package main

import (
	"strings"
)

type HexColor struct {
	code string
}

func NewHexColor(code string) HexColor {
	return HexColor{code: strings.ToLower(code)}
}
