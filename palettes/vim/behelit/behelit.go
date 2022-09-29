package behelit

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x00ff87,
		0x121212,
		0x1c1c1c,
		0x262626,
		0x4e4e4e,
		0x5f5f87,
		0x5f5faf,
		0x5f87ff,
		0x5fff5f,
		0xd70057,
		0xd7005f,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
