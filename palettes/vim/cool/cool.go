package cool

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f87,
		0x008787,
		0x0087af,
		0x00afa2,
		0x2e8700,
		0x324e59,
		0x466d79,
		0x47af00,
		0x585858,
		0x5f005f,
		0x872800,
		0x875300,
		0x875faf,
		0xaf0000,
		0xaf2800,
		0xaf5f00,
		0xd78700,
		0xe4e4e4,
		0xeeeeee,
		0xff0000,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
