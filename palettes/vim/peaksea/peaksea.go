package peaksea

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
		0x3a3a3a,
		0x4e4e4e,
		0x585858,
		0x60f080,
		0x666666,
		0xc0d8f8,
		0xd0d090,
		0xd75f5f,
		0xe0c060,
		0xe4e4e4,
		0xeeeeee,
		0xf0c0f0,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
