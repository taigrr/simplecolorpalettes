package desertink

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x004866,
		0x005f00,
		0x005f87,
		0x0087af,
		0x080808,
		0x303030,
		0x3a3a3a,
		0x444444,
		0x4a4a4a,
		0x777777,
		0x870000,
		0x999999,
		0x99ddff,
		0xafd700,
		0xb2e5ff,
		0xbbbbbb,
		0xd74444,
		0xff8700,
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
