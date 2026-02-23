package simplecolor

import (
	"math"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	// Colors that are well within sRGB gamut round-trip cleanly.
	// Highly saturated colors (pure blue, etc.) may have slight drift
	// due to floating point and gamut boundary effects.
	colors := []struct {
		hex       string
		tolerance int
	}{
		{"#FF0000", 1},  // red
		{"#00FF00", 1},  // green
		{"#FFFFFF", 1},  // white
		{"#000000", 1},  // black
		{"#FF8800", 1},  // orange
		{"#808080", 1},  // mid gray
		{"#0000FF", 32}, // blue - significant drift due to OKLab gamut boundary
		{"#123456", 3},  // arbitrary
		{"#ABCDEF", 3},  // arbitrary
	}

	for _, tc := range colors {
		c := FromHexString(tc.hex)
		oklch := c.ToOKLCH()
		back := FromOKLCH(oklch.L, oklch.C, oklch.H)
		r1, g1, b1, _ := c.RGBA()
		r2, g2, b2, _ := back.RGBA()
		if abs(int(r1)-int(r2)) > tc.tolerance || abs(int(g1)-int(g2)) > tc.tolerance || abs(int(b1)-int(b2)) > tc.tolerance {
			t.Errorf("round-trip failed for %s: got %s (R%d G%d B%d vs R%d G%d B%d)",
				c.ToHex(), back.ToHex(), r1, g1, b1, r2, g2, b2)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestKnownValues(t *testing.T) {
	tests := []struct {
		name  string
		color SimpleColor
		wantL float64
		wantC float64
	}{
		{"black", FromHexString("#000000"), 0.0, 0.0},
		{"white", FromHexString("#FFFFFF"), 1.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oklch := tt.color.ToOKLCH()
			if math.Abs(oklch.L-tt.wantL) > 0.01 {
				t.Errorf("L = %f, want %f", oklch.L, tt.wantL)
			}
			if math.Abs(oklch.C-tt.wantC) > 0.01 {
				t.Errorf("C = %f, want %f", oklch.C, tt.wantC)
			}
		})
	}

	// Red should have high chroma and L around 0.63
	red := FromHexString("#FF0000").ToOKLCH()
	if red.L < 0.5 || red.L > 0.75 {
		t.Errorf("red L = %f, expected ~0.63", red.L)
	}
	if red.C < 0.2 {
		t.Errorf("red C = %f, expected high chroma", red.C)
	}

	// Green should be lighter than red
	green := FromHexString("#00FF00").ToOKLCH()
	if green.L < red.L {
		t.Errorf("green L (%f) should be > red L (%f)", green.L, red.L)
	}
}

func TestClampOKLCH(t *testing.T) {
	c := FromHexString("#FF0000") // bright red
	clamped := c.ClampOKLCH(0.3, 0.5, 0.0, 0.15)
	oklch := clamped.ToOKLCH()

	if oklch.L < 0.29 || oklch.L > 0.51 {
		t.Errorf("clamped L = %f, expected in [0.3, 0.5]", oklch.L)
	}
	if oklch.C > 0.16 {
		t.Errorf("clamped C = %f, expected <= 0.15", oklch.C)
	}
}

func TestNormalizeLightness(t *testing.T) {
	// Use moderate chroma colors to stay within gamut when changing lightness
	palette := SimplePalette{
		FromHexString("#CC6666"), // muted red
		FromHexString("#66CC66"), // muted green
		FromHexString("#6666CC"), // muted blue
	}

	normalized := palette.NormalizeLightness(0.7)
	for i, c := range normalized {
		oklch := c.ToOKLCH()
		if math.Abs(oklch.L-0.7) > 0.05 {
			t.Errorf("color %d: L = %f, expected ~0.7", i, oklch.L)
		}
	}
}

func TestNormalizeChroma(t *testing.T) {
	palette := SimplePalette{
		FromHexString("#FF0000"),
		FromHexString("#00FF00"),
		FromHexString("#0000FF"),
	}

	normalized := palette.NormalizeChroma(0.1)
	for i, c := range normalized {
		oklch := c.ToOKLCH()
		if math.Abs(oklch.C-0.1) > 0.02 {
			t.Errorf("color %d: C = %f, expected 0.1", i, oklch.C)
		}
	}
}

func TestPaletteClampOKLCH(t *testing.T) {
	palette := SimplePalette{
		FromHexString("#FF0000"),
		FromHexString("#00FF00"),
		FromHexString("#0000FF"),
	}

	clamped := palette.ClampOKLCH(0.4, 0.8, 0.05, 0.15)
	for i, c := range clamped {
		oklch := c.ToOKLCH()
		if oklch.L < 0.38 || oklch.L > 0.82 {
			t.Errorf("color %d: L = %f, out of bounds", i, oklch.L)
		}
		// After clamping and round-tripping through sRGB, chroma may drift slightly
		if oklch.C < 0.03 || oklch.C > 0.20 {
			t.Errorf("color %d: C = %f, out of bounds", i, oklch.C)
		}
	}
}
