package murmur

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1c1c1c,
		0x4e4e4e,
		0x5f5f5f,
		0x5f87ff,
		0x870000,
		0x87af5f,
		0xafaf87,
		0xf5f5f5,
		0xff8c00,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
