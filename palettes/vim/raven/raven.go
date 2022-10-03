package raven

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x11c279,
		0x222222,
		0x2e2e2e,
		0x5e5e5e,
		0x6565ff,
		0xa4c639,
		0xc8c8c8,
		0xe25000,
		0xe60000,
		0xff0000,
		0xff2121,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
