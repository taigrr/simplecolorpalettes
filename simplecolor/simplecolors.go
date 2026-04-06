package simplecolor

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"
	"strings"
)

const TotalHexColorspace = 0xffffff + 1

type (
	SimpleColor   int
	NamedPalette  map[string]SimpleColor
	SimplePalette []SimpleColor
)

func (n NamedPalette) Get(name string) SimpleColor {
	return n[name]
}

func (n NamedPalette) Names() []string {
	names := make([]string, 0, len(n))
	for k := range n {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

func (n NamedPalette) ToSimplePalette() SimplePalette {
	var sp SimplePalette
	for _, c := range n {
		sp = append(sp, c)
	}
	return sp
}

func (n NamedPalette) ToPalette() color.Palette {
	var x color.Palette
	for _, c := range n {
		x = append(x, SimpleColor(c))
	}
	return color.Palette(x)
}

func (s SimplePalette) Sort() SimplePalette {
	sort.Sort(s)
	return s
}

func (s SimplePalette) Join(b SimplePalette) SimplePalette {
	m := make(map[SimpleColor]SimpleColor)
	r := SimplePalette{}
	for _, c := range s {
		m[c] = c
	}
	for _, c := range b {
		m[c] = c
	}
	for _, c := range m {
		r = append(r, c)
	}
	sort.Sort(r)
	return r
}

func (s SimplePalette) Get(index int) SimpleColor {
	if index < 0 || index >= len(s) {
		return FromHexString("#66042d")
	}
	return s[index]
}

func (s SimplePalette) ToPalette() color.Palette {
	var x color.Palette
	for _, c := range s {
		x = append(x, SimpleColor(c))
	}
	return color.Palette(x)
}

func (c SimpleColor) ToHex() string {
	return "#" + fmt.Sprintf("%06X", c)
}

// ToShortHex returns a 3-character hex string if possible (e.g. "#F00"),
// otherwise falls back to the full 6-character hex string.
func (c SimpleColor) ToShortHex() string {
	r := uint32(c) >> 16 & 0xFF
	g := uint32(c) >> 8 & 0xFF
	b := uint32(c) & 0xFF
	// A color can be shortened when each channel's high and low nibbles match
	if r>>4 == r&0xF && g>>4 == g&0xF && b>>4 == b&0xF {
		return "#" + fmt.Sprintf("%X%X%X", r&0xF, g&0xF, b&0xF)
	}
	return c.ToHex()
}

// RGBA implements the color.Color interface. Returns pre-multiplied 16-bit
// values in [0, 0xFFFF] as required by the standard library.
func (c SimpleColor) RGBA() (r, g, b, a uint32) {
	r8 := uint32(c) >> 16 & 0xFF
	g8 := uint32(c) >> 8 & 0xFF
	b8 := uint32(c) & 0xFF
	return r8 | r8<<8, g8 | g8<<8, b8 | b8<<8, 0xFFFF
}

// RGB8 returns the 8-bit [0, 255] red, green, and blue components.
func (c SimpleColor) RGB8() (r, g, b uint8) {
	return uint8(uint32(c) >> 16 & 0xFF), uint8(uint32(c) >> 8 & 0xFF), uint8(uint32(c) & 0xFF)
}

func (c SimpleColor) ToAnsi16() SimpleColor {
	converted := ansi[0:16].ToPalette().Convert(c)
	r, g, b, _ := converted.RGBA()
	return SimpleColor((r>>8)<<16 | (g>>8)<<8 | (b >> 8))
}

func (c SimpleColor) ToExtendedAnsi() SimpleColor {
	converted := ansi.ToPalette().Convert(c)
	r, g, b, _ := converted.RGBA()
	return SimpleColor((r>>8)<<16 | (g>>8)<<8 | (b >> 8))
}

func (n NamedPalette) ToExtendedAnsi() NamedPalette {
	for k, v := range n {
		n[k] = v.ToExtendedAnsi()
	}
	return n
}

func (s SimplePalette) ToExtendedAnsi() (sp SimplePalette) {
	used := make(map[SimpleColor]bool)
	for _, x := range s {
		clampedColor := x.ToExtendedAnsi()
		if _, found := used[clampedColor]; found {
			continue
		}
		used[clampedColor] = true
		sp = append(sp, x.ToExtendedAnsi())
	}
	sort.Sort(sp)

	return
}

func (s SimplePalette) Len() int {
	return len(s)
}

func (s SimplePalette) Less(i, j int) bool {
	return int(s[i]) < int(s[j])
}

func (s SimplePalette) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SimplePalette) ToAnsi16() (sp SimplePalette) {
	used := make(map[SimpleColor]bool)
	for _, x := range s {
		clampedColor := x.ToAnsi16()
		if _, found := used[clampedColor]; found {
			continue
		}
		used[clampedColor] = true
		sp = append(sp, x.ToExtendedAnsi())
	}
	sort.Sort(sp)

	return
}

func New(color int) SimpleColor {
	return SimpleColor(color % TotalHexColorspace)
}

// FromRGBA creates a SimpleColor from 8-bit [0, 255] R, G, B values.
// The alpha component is ignored.
func FromRGBA(r, g, b, _ uint32) SimpleColor {
	return SimpleColor((r&0xFF)<<16 | (g&0xFF)<<8 | (b & 0xFF))
}

func FromHexString(h string) SimpleColor {
	h = strings.ReplaceAll(h, "#", "")
	hexRunes := []rune(h)
	switch len(hexRunes) {
	case 6:
		break
	case 3:
		stretchedHex := fmt.Sprintf("%s%s%s%s%s%s",
			string(hexRunes[0]),
			string(hexRunes[0]),
			string(hexRunes[1]),
			string(hexRunes[1]),
			string(hexRunes[2]),
			string(hexRunes[2]))
		h = stretchedHex
	default:
		return FromHexString("#66042d")
	}
	i, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		// if there's an error, we can't return it
		// since this is a simple method.
		// Instead, return a really weird color to draw attention.
		// Black is the zero-value and it's actually used often, "plum" is
		// (probably) not.
		return FromHexString("#66042d")
	}
	return SimpleColor(i)
}
