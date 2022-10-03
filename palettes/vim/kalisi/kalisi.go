package kalisi

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x005faf,
		0x0087ff,
		0x204d20,
		0x404042,
		0x5f0000,
		0x5f005f,
		0x5fafff,
		0x8700af,
		0x87d7ff,
		0xa6db29,
		0xafd700,
		0xbcbcbc,
		0xc5c5c5,
		0xd75fff,
		0xd7ff00,
		0xe80000,
		0xff0000,
		0xff87ff,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
