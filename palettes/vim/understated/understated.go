package understated

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x080808,
		0x1c1c1c,
		0x262626,
		0x303030,
		0x4e4e4e,
		0x5f005f,
		0x5f5f5f,
		0x5f87ff,
		0x870000,
		0x87af5f,
		0xaf5f00,
		0xafaf87,
		0xff0000,
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
