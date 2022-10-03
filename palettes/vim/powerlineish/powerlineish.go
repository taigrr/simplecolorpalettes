package powerlineish

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x005f5f,
		0x005f87,
		0x0087af,
		0x080808,
		0x121212,
		0x303030,
		0x5fafd7,
		0x87d7ff,
		0x9e9e9e,
		0xafd700,
		0xd70000,
		0xffaf00,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
