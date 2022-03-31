package simplecolor

import (
	"image/color"
	"sort"
)

type SimpleColor int
type NamedPalette map[string]SimpleColor
type SimplePalette []SimpleColor
type conversionPalette []color.Color

func (s SimplePalette) ToPalette() color.Palette {
	var x color.Palette
	for _, c := range s {
		x = append(x, SimpleColor(c))
	}
	return color.Palette(x)
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
