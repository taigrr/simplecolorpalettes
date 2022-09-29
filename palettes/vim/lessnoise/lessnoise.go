package lessnoise

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x080808,
		0x121212,
		0x1c1c1c,
		0x303030,
		0x444444,
		0x5f87ff,
		0x6c6c6c,
		0x87afff,
		0x9e9e9e,
		0xafd7ff,
		0xafffff,
		0xb2b2b2,
		0xd0d0d0,
		0xd75f5f,
		0xeeeeee,
		0xff5f5f,
		0xffafd7,
		0xffffaf,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
