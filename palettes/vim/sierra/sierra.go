package sierra

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x2a2a2a,
		0x303030,
		0x545454,
		0x666666,
		0x686868,
		0xaf5f5f,
		0xaf8787,
		0xafd7d7,
		0xd75f5f,
		0xdfaf87,
		0xf7e4c0,
		0xffafaf,
		0xffb2af,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
