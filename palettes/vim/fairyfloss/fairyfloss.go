package fairyfloss

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x3b3a32,
		0x49483e,
		0x63588d,
		0x8076aa,
		0x80ffbd,
		0xae81ff,
		0xc2ffdf,
		0xc5a3ff,
		0xe6c000,
		0xefe6ff,
		0xf8f8f0,
		0xf92672,
		0xff857f,
		0xffb8d1,
		0xfffea0,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
