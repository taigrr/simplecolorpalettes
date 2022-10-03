package minimalist

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x1c1c1c,
		0x262626,
		0x3a3a3a,
		0x4e4e4e,
		0x666666,
		0xd75f5f,
		0xe4e4e4,
		0xeeeeee,
		0xffaf5f,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
