package main

import "testing"

func TestVerify(t *testing.T) {
	hex3 := NewHexColor("012")
	hex4 := NewHexColor("abcd")
	hex6 := NewHexColor("987fed")
	hex12 := NewHexColor("01ab23cd67ef")
	nonhex := NewHexColor("0g9h1i")

	if hex3.Verify() != nil {
		t.Fatal("hex3.Verify(): should be succeed, but failed")
	}
	if hex6.Verify() != nil {
		t.Fatal("hex6.Verify(): should be succeed, but failed")
	}
	if hex4.Verify() == nil {
		t.Fatal("hex4.Verify(): should be failed, but succeed")
	}
	if hex12.Verify() == nil {
		t.Fatal("hex12.Verify(): should be failed, but succeed")
	}
	if nonhex.Verify() == nil {
		t.Fatal("nonhex.Verify(): should be failed, but succeed")
	}
}
