package distinguished

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x000000,
		0x1c1c1c,
		0x444444,
		0x5f87af,
		0x8a8a8a,
		0xaf5f5f,
		0xafaf5f,
		0xbcbcbc,
		0xd70000,
		0xd75f00,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
