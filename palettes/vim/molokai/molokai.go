package molokai

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x080808,
		0x1b1d1e,
		0x232526,
		0x465457,
		0x66d9ef,
		0xa6e22e,
		0xe6db74,
		0xf8f8f0,
		0xf92672,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
