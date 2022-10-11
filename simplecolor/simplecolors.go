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
	SimpleColor       int
	NamedPalette      map[string]SimpleColor
	SimplePalette     []SimpleColor
	conversionPalette []color.Color
)

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
	if index < 0 || index > len(s) {
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

func (c SimpleColor) ToShortHex() string {
	value := c >> 16 & 0xF
	value += c >> 8 & 0xF
	value += c & 0xF
	return "#" + fmt.Sprintf("%06X", value)
}

func (c SimpleColor) RGBA() (r, g, b, a uint32) {
	return uint32(c) >> 16 & 0xFF, uint32(c) >> 8 & 0xFF, uint32(c) & 0xFF, 0xFF
}

func (c SimpleColor) ToAnsi16() SimpleColor {
	color := ansi[0:16].ToPalette().Convert(c)
	r, g, b, _ := color.RGBA()
	return SimpleColor(uint32(r)<<16 + uint32(g)<<8 + b)
}

func (c SimpleColor) ToExtendedAnsi() SimpleColor {
	color := ansi.ToPalette().Convert(c)
	r, g, b, _ := color.RGBA()
	return SimpleColor(uint32(r)<<16 + uint32(g)<<8 + b)
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

func FromRGBA(r, g, b, _ uint32) SimpleColor {
	c := r % TotalHexColorspace
	c = c<<8 + (g % TotalHexColorspace)
	c = c<<8 + (b % TotalHexColorspace)
	return SimpleColor(c)
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
