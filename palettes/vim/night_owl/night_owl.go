package night_owl

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x282c34,
		0x6788cc,
		0x68b0a0,
		0x81aaff,
		0x83dcc8,
		0x8cac4c,
		0x9f74bb,
		0xafd75f,
		0xc792ea,
		0xccac6c,
		0xffd787,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
