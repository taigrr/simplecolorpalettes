package simplecolor

import (
	"testing"
)

func TestCreateColor(t *testing.T) {
	tc := []struct {
		ID     string
		Value  int
		Result int
	}{
		{ID: "white", Value: TotalHexColorspace - 1, Result: TotalHexColorspace - 1},
		{ID: "WraparoundBlack", Value: TotalHexColorspace, Result: 0},
		{ID: "black", Value: 0, Result: 0},
	}
	for _, c := range tc {
		t.Run(c.ID, func(t *testing.T) {
			color := New(c.Value)
			if int(color) != c.Result {
				t.Errorf("Expected Value %d, but got %d", c.Result, color)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	tc := []struct {
		ID     string
		R      uint32
		G      uint32
		B      uint32
		Result string
	}{
		{ID: "red", R: 0xFF, G: 0x00, B: 0x00, Result: "#FF0000"},
		{ID: "black", R: 0, G: 0, B: 0, Result: "#000000"},
	}
	for _, c := range tc {
		t.Run(c.ID, func(t *testing.T) {
			color := FromRGBA(c.R, c.G, c.B, 0)
			if color.ToHex() != c.Result {
				t.Errorf("Expected Value %s, but got %s", c.Result, color.ToHex())
			}
		})
	}
}

func TestRGB(t *testing.T) {
	tc := []struct {
		ID     string
		R      uint32
		G      uint32
		B      uint32
		Result int
	}{
		{ID: "white", R: 255, G: 255, B: 255, Result: TotalHexColorspace - 1},
		{ID: "black", R: 0, G: 0, B: 0, Result: 0},
	}
	for _, c := range tc {
		t.Run(c.ID, func(t *testing.T) {
			color := FromRGBA(c.R, c.G, c.B, 0)
			if int(color) != c.Result {
				t.Errorf("Expected Value %d, but got %d", c.Result, color)
			}
		})
	}
}

func TestFromHex(t *testing.T) {
	tc := []struct {
		ID     string
		Input  string
		Result string
		Error  error
	}{
		{ID: "redshort", Input: "#F00", Result: "#FF0000", Error: nil},
		{ID: "redlong", Input: "#F00000", Result: "#F00000", Error: nil},
	}
	for _, c := range tc {
		t.Run(c.ID, func(t *testing.T) {
			color := FromHexString(c.Input)
			if color.ToHex() != c.Result {
				t.Errorf("Expected Value %s, but got %s", c.Result, color.ToHex())
			}
		})
	}
}
