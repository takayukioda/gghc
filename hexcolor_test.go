package main

import (
	"fmt"
	"testing"
)

func failure(it string, expected string, actual string) string {
	return fmt.Sprintf("%s: expected to be %s, got %s", it, expected, actual)
}

func TestNew(t *testing.T) {
	hex, lower := NewHexColor("aB1FE0"), "ab1fe0"
	if hex.GetCode() != lower {
		t.Fatal(failure("NewHexColor", lower, hex.GetCode()))
	}
}

func TestVerify(t *testing.T) {
	hex3 := NewHexColor("012")
	hex4 := NewHexColor("abcd")
	hex6 := NewHexColor("987fed")
	hex12 := NewHexColor("01ab23cd67ef")
	nonhex := NewHexColor("0g9h1i")

	if hex3.Verify() != nil {
		t.Fatal("length 3, Verify: expected to be nil, got", hex3.Verify())
	}
	if hex6.Verify() != nil {
		t.Fatal("length 6, Verify: expected to be nil, got", hex6.Verify())
	}
	if hex4.Verify() == nil {
		t.Fatal("length 4, Verify: expected to be an error, got nil")
	}
	if hex12.Verify() == nil {
		t.Fatal("length 12, Verify: expected to be an error, got nil")
	}
	if nonhex.Verify() == nil {
		t.Fatal("non hex, Verify: expected to be an error, got nil")
	}
}

func TestToFull(t *testing.T) {
	hex3, hex3full := NewHexColor("f12"), "ff1122"
	hex6 := NewHexColor("d97f8e")

	if hex3.ToFull() != hex3full {
		t.Fatal(failure("length 3, ToFull", hex3full, hex3.ToFull()))
	}
	if hex6.ToFull() != hex6.GetCode() {
		t.Fatal(failure("length 3, ToFull", hex3full, hex3.ToFull()))
	}
}

func TestCompress(t *testing.T) {
	hex3, hex3comp := NewHexColor("9a3"), "9a3"
	hex6, hex6comp := NewHexColor("ee8833"), "e83"
	hexe := NewHexColor("d97f8e")

	if comp, err := hex3.Compress(); err != nil {
		t.Fatal(hex3.GetCode(), "Compress(): should not be an error")
	} else if comp != hex3comp {
		t.Fatal(hex3.GetCode(), "Compress(): should equal to", hex3comp)
	}
	if comp, err := hex6.Compress(); err != nil {
		t.Fatal(hex6.GetCode(), "Compress(): should not be an error")
	} else if comp != hex6comp {
		t.Fatal(hex6.GetCode(), "Compress(): should equal to", hex6comp)
	}
	if _, err := hexe.Compress(); err == nil {
		t.Fatal(hexe.GetCode(), "Compress(): should be an error")
	}
}
