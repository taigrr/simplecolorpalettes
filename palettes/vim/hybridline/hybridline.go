package hybridline

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x000000,
		0x005f00,
		0x005f5f,
		0x005f87,
		0x0087af,
		0x1c1c1c,
		0x262626,
		0x282a2e,
		0x303030,
		0x373b41,
		0x4e4e4e,
		0x8abeb7,
		0xac4142,
		0xb5bd68,
		0xc5c8c6,
		0xcc6666,
		0xde935f,
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
