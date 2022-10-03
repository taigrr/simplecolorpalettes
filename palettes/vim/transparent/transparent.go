package transparent

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1d1f21,
		0x2f3d4d,
		0x3f4b59,
		0x8d96a1,
		0xbbe67e,
		0xd4bfff,
		0xd70000,
		0xf07178,
		0xffae57,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
