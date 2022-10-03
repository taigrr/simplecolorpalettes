package cobalt2

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1780e9,
		0x1a3548,
		0x1f7ad8,
		0x204458,
		0x46dd3c,
		0x4e4e4e,
		0x666d51,
		0x8cc2fd,
		0xb42839,
		0xea9299,
		0xfee533,
		0xff9d00,
		0xffff9a,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
