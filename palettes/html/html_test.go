package html

import (
	"slices"
	"testing"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

func TestGetNamedPalette(t *testing.T) {
	namedPalette := GetNamedPalette()

	tests := []struct {
		name string
		want simplecolor.SimpleColor
	}{
		{name: "darkgoldenrod", want: simplecolor.FromHexString("#B8860B")},
		{name: "gray", want: simplecolor.FromHexString("#808080")},
		{name: "khaki", want: simplecolor.FromHexString("#F0E68C")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := namedPalette.Get(test.name); got != test.want {
				t.Fatalf("Get(%q) = %s, want %s", test.name, got.ToHex(), test.want.ToHex())
			}
		})
	}
}

func TestGetPalette(t *testing.T) {
	palette := GetPalette()
	namedPalette := GetNamedPalette()

	if len(palette) != len(namedPalette) {
		t.Fatalf("GetPalette() len = %d, want %d", len(palette), len(namedPalette))
	}
	if !slices.IsSorted(palette) {
		t.Fatal("GetPalette() should return sorted colors")
	}
}
