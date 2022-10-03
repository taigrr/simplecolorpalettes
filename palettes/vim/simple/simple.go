package simple

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x00dfff,
		0x080808,
		0x1c1c1c,
		0x4e4e4e,
		0x5fff00,
		0x767676,
		0xaf0000,
		0xd78700,
		0xdf0000,
		0xdfdf00,
		0xff5f00,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
