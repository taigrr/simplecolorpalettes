package light

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x000087,
		0x005f00,
		0x005f5f,
		0x005fff,
		0x00875f,
		0x00df87,
		0x00dfff,
		0x5f0000,
		0x666666,
		0x8a8a8a,
		0xa8a8a8,
		0xafff87,
		0xafffff,
		0xb2b2b2,
		0xd0d0d0,
		0xd78700,
		0xdf0000,
		0xdf5f00,
		0xff0000,
		0xff5f00,
		0xffaf00,
		0xffdfdf,
		0xffff87,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
