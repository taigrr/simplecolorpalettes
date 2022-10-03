package supernova

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1d2529,
		0x242e33,
		0x585858,
		0x585864,
		0xd40059,
		0xd7005f,
		0xe4e4e4,
		0xeeeeee,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
