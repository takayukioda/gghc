package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	color, _ := NewHexColor("aB1FE0")
	if color.GetCode() != "ab1fe0" {
		t.Fatalf("NewHexColor: expected initialized code to be lower case, got %s",
			color.GetCode())
	}
}

func TestVerify(t *testing.T) {
	if _, errs := NewHexColor("012"); errs != nil {
		t.Fatalf("Verify: 3-chars color code should succeed, got %s", errs)
	}
	if _, errs := NewHexColor("987fed"); errs != nil {
		t.Fatalf("Verify: 6-chars color code should succeed, got %s", errs)
	}
	if color, errs := NewHexColor("abcd"); errs == nil {
		t.Fatalf("Verify: 4-chars color code should fail, got %s", color)
	}
	if color, errs := NewHexColor("01ab23cd67ef"); errs == nil {
		t.Fatalf("Verify: 12-chars color code should fail, got %s", color)
	}
	if color, errs := NewHexColor("0g9h1i"); errs == nil {
		t.Fatalf("Verify: code with non-hex character should fail, got %s", color)
	}
}

func TestToFull(t *testing.T) {
	if color, _ := NewHexColor("f12"); color.ToFull() != "ff1122" {
		t.Fatalf("ToFull: 3-chars color code should be extened, got %s", color.ToFull())
	}
	if color, _ := NewHexColor("d97f8e"); color.ToFull() != color.GetCode() {
		t.Fatalf("ToFull: 6-chars color code should equals to original code, got %s",
			color.ToFull())
	}
}

func TestCompress(t *testing.T) {
	var color *HexColor
	var _ []error

	color, _ = NewHexColor("d97f8e")
	if comp, err := color.Compress(); err == nil {
		t.Fatalf("Compress: color code with non-supported format should fail, got %s", comp)
	}

	color, _ = NewHexColor("9a3")
	if comp, _ := color.Compress(); comp != color.GetCode() {
		t.Fatalf("Compress: 6-chars color code should compressed, got %s", comp)
	}

	color, _ = NewHexColor("ee8833")
	if comp, _ := color.Compress(); comp != "e83" {
		t.Fatalf("Compress: 3-chars color code should euqals to original code, got %s", comp)
	}
}
