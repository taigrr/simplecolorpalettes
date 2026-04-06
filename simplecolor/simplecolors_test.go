package simplecolor

import (
	"image/color"
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

func TestColorInterface(t *testing.T) {
	// Verify SimpleColor implements color.Color with proper 16-bit values
	var _ color.Color = SimpleColor(0)

	red := FromHexString("#FF0000")
	r, g, b, a := red.RGBA()
	if r != 0xFFFF {
		t.Errorf("red R: got %d, want %d", r, 0xFFFF)
	}
	if g != 0 {
		t.Errorf("red G: got %d, want 0", g)
	}
	if b != 0 {
		t.Errorf("red B: got %d, want 0", b)
	}
	if a != 0xFFFF {
		t.Errorf("red A: got %d, want %d", a, 0xFFFF)
	}

	// Test a mid-value color
	mid := FromHexString("#808080")
	mr, mg, mb, _ := mid.RGBA()
	if mr != 0x8080 || mg != 0x8080 || mb != 0x8080 {
		t.Errorf("mid RGBA: got (0x%04X, 0x%04X, 0x%04X), want all 0x8080", mr, mg, mb)
	}
}

func TestRGB8(t *testing.T) {
	c := FromHexString("#1A2B3C")
	r, g, b := c.RGB8()
	if r != 0x1A || g != 0x2B || b != 0x3C {
		t.Errorf("RGB8: got (%d, %d, %d), want (26, 43, 60)", r, g, b)
	}
}

func TestToShortHex(t *testing.T) {
	tc := []struct {
		input string
		want  string
	}{
		{"#FF0000", "#F00"},
		{"#00FF00", "#0F0"},
		{"#AABBCC", "#ABC"},
		{"#000000", "#000"},
		{"#FFFFFF", "#FFF"},
		{"#123456", "#123456"}, // cannot be shortened
		{"#F00000", "#F00000"}, // 0xF0 high!=low nibble
	}
	for _, c := range tc {
		got := FromHexString(c.input).ToShortHex()
		if got != c.want {
			t.Errorf("ToShortHex(%s): got %s, want %s", c.input, got, c.want)
		}
	}
}

func TestGetBounds(t *testing.T) {
	p := SimplePalette{FromHexString("#FF0000"), FromHexString("#00FF00")}
	// Valid indices
	if p.Get(0).ToHex() != "#FF0000" {
		t.Error("Get(0) failed")
	}
	if p.Get(1).ToHex() != "#00FF00" {
		t.Error("Get(1) failed")
	}
	// Out of bounds should return fallback
	fallback := FromHexString("#66042d").ToHex()
	if p.Get(-1).ToHex() != fallback {
		t.Error("Get(-1) should return fallback")
	}
	if p.Get(2).ToHex() != fallback {
		t.Error("Get(len) should return fallback")
	}
}

func TestNamedPaletteNamesSorted(t *testing.T) {
	np := NamedPalette{
		"zebra":    FromHexString("#000000"),
		"apple":    FromHexString("#FF0000"),
		"mango":    FromHexString("#FFAA00"),
		"bluebird": FromHexString("#0000FF"),
	}
	names := np.Names()
	for i := 1; i < len(names); i++ {
		if names[i] < names[i-1] {
			t.Errorf("Names() not sorted: %v", names)
			break
		}
	}
}

func TestPaletteConvert(t *testing.T) {
	// color.Palette.Convert uses Euclidean distance in the RGBA space.
	// With proper 16-bit RGBA, this should find the nearest color correctly.
	p := SimplePalette{
		FromHexString("#000000"),
		FromHexString("#FFFFFF"),
	}.ToPalette()

	// A dark color should map to black
	nearest := p.Convert(FromHexString("#111111"))
	r, _, _, _ := nearest.RGBA()
	if r != 0 {
		t.Errorf("expected dark color to map to black, got R=%d", r)
	}

	// A light color should map to white
	nearest = p.Convert(FromHexString("#EEEEEE"))
	rW, _, _, _ := nearest.RGBA()
	if rW != 0xFFFF {
		t.Errorf("expected light color to map to white, got R=%d", rW)
	}
}
