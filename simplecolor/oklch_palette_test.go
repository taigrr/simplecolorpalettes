package simplecolor

import (
	"slices"
	"testing"
)

func TestOKLCHPaletteTransformsPreserveInput(t *testing.T) {
	// Unsorted colors and a duplicate make order and multiplicity observable.
	original := SimplePalette{0xCC9966, 0x336699, 0xCC9966, 0x669966}
	tests := []struct {
		name      string
		transform func(SimplePalette) SimplePalette
		want      SimplePalette
	}{
		{
			name: "clamp",
			transform: func(palette SimplePalette) SimplePalette {
				return palette.ClampOKLCH(0.5, 0.7, 0.02, 0.08)
			},
			want: SimplePalette{
				original[0].ClampOKLCH(0.5, 0.7, 0.02, 0.08),
				original[1].ClampOKLCH(0.5, 0.7, 0.02, 0.08),
				original[2].ClampOKLCH(0.5, 0.7, 0.02, 0.08),
				original[3].ClampOKLCH(0.5, 0.7, 0.02, 0.08),
			},
		},
		{
			name: "lightness",
			transform: func(palette SimplePalette) SimplePalette {
				return palette.NormalizeLightness(0.6)
			},
		},
		{
			name: "chroma",
			transform: func(palette SimplePalette) SimplePalette {
				return palette.NormalizeChroma(0.05)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			if want == nil {
				for _, color := range original {
					value := color.ToOKLCH()
					if test.name == "lightness" {
						value.L = 0.6
					} else {
						value.C = 0.05
					}
					want = append(want, FromOKLCH(value.L, value.C, value.H))
				}
			}
			input := slices.Clone(original)
			got := test.transform(input)
			if !slices.Equal(got, want) {
				t.Fatalf("transform = %v, want %v (same order and duplicates)", got, want)
			}
			if !slices.Equal(input, original) {
				t.Fatal("transform modified its input")
			}
			got[0] = 0
			if !slices.Equal(input, original) {
				t.Fatal("result aliases its input")
			}
			for _, empty := range []SimplePalette{nil, {}} {
				if got := test.transform(empty); len(got) != 0 {
					t.Errorf("empty input produced %v", got)
				}
			}
		})
	}
}
