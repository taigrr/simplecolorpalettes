package simplecolor

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"
	"strings"
)

type (
	SimpleColor       int
	NamedPalette      map[string]SimpleColor
	SimplePalette     []SimpleColor
	conversionPalette []color.Color
)

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

func (s SimpleColor) RGBA() (r, g, b, a uint32) {
	return uint32(s) >> 16 & 0xFF, uint32(s) >> 8 & 0xFF, uint32(s) & 0xFF, 0xFF
}

func (input SimpleColor) ToAnsi16() SimpleColor {
	color := ansi[0:16].ToPalette().Convert(input)
	r, g, b, _ := color.RGBA()
	return SimpleColor(uint32(r)<<16 + uint32(g)<<8 + b)
}

func (input SimpleColor) ToExtendedAnsi() SimpleColor {
	color := ansi.ToPalette().Convert(input)
	r, g, b, _ := color.RGBA()
	return SimpleColor(uint32(r)<<16 + uint32(g)<<8 + b)
}

func (p SimplePalette) ToExtendedAnsi() (sp SimplePalette) {
	used := make(map[SimpleColor]bool)
	for _, x := range p {
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

func (e SimplePalette) Len() int {
	return len(e)
}

func (e SimplePalette) Less(i, j int) bool {
	return int(e[i]) < int(e[j])
}

func (e SimplePalette) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (p SimplePalette) ToAnsi16() (sp SimplePalette) {
	used := make(map[SimpleColor]bool)
	for _, x := range p {
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
	return SimpleColor(color)
}

func FromRGBA(r, g, b, _ uint32) SimpleColor {
	c := r
	c = c<<8 + g
	c = c<<8 + b
	return SimpleColor(c)
}

func FromHexString(h string) SimpleColor {
	h = strings.ReplaceAll(h, "#", "")
	hexRunes := []rune(h)
	switch len(hexRunes) {
	case 6:
		break
	case 3:
		stretchedHex := hexRunes[0] + hexRunes[0] + hexRunes[1] + hexRunes[1] + hexRunes[2] + hexRunes[2]
		h = string(stretchedHex)
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
