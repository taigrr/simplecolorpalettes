package apprentice

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1c1c1c,
		0x262626,
		0x585858,
		0x5f875f,
		0x87875f,
		0x8787af,
		0x87af87,
		0x87afd7,
		0xaf5f5f,
		0xbcbcbc,
		0xff8700,
		0xffffaf,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
