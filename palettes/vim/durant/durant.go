package durant

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x005f87,
		0x00875f,
		0x073642,
		0x1a1a18,
		0x1c1c1c,
		0x262626,
		0x2e2d2a,
		0x303030,
		0x44403a,
		0x4e4e4e,
		0x586e75,
		0x5f005f,
		0x5f00af,
		0x875faf,
		0x875fd7,
		0x87d7ff,
		0x90a680,
		0x93a1a1,
		0x9e9e9e,
		0xaf0000,
		0xafd700,
		0xd78700,
		0xd7d7ff,
		0xff0000,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
