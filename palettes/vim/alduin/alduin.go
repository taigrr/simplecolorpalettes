package alduin

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1c1c1c,
		0x2a2a2a,
		0x545454,
		0x5f5f87,
		0x666666,
		0x875f5f,
		0x87875f,
		0x878787,
		0x87afaf,
		0xaf1600,
		0xaf5f00,
		0xaf5f5f,
		0xaf875f,
		0xaf8787,
		0xafd7d7,
		0xdfaf87,
		0xdfdfaf,
		0xffdf87,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
