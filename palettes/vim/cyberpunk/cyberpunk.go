package cyberpunk

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0197dd,
		0x0c35bf,
		0x0eeafa,
		0x191919,
		0x408000,
		0x414c3b,
		0x666666,
		0x800080,
		0x87f025,
		0xcdb1ad,
		0xfe1bf0,
		0xff0000,
		0xffd302,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
