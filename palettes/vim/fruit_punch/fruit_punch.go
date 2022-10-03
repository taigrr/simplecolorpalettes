package fruit_punch

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x303030,
		0x3a3a3a,
		0x444444,
		0x79e5e0,
		0xb2b2b2,
		0xe8a15a,
		0xf29db4,
		0xf97070,
		0xfce78d,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
