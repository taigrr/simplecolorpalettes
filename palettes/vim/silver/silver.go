package silver

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0000b3,
		0x007599,
		0x0d935c,
		0x414141,
		0xa1a1a1,
		0xb30000,
		0xdddddd,
		0xe1e1e1,
		0xe25000,
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
