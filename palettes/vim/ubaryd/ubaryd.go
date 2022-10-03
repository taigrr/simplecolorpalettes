package ubaryd

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x242321,
		0x416389,
		0x666462,
		0x9a4820,
		0xb88853,
		0xc14c3d,
		0xe25a74,
		0xf4cf86,
		0xf8f6f2,
		0xf9ef6d,
		0xff7400,
		0xffa724,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
