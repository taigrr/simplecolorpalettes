package bubblegum

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
		0x87afd7,
		0xafafd7,
		0xafd787,
		0xb2b2b2,
		0xd78787,
		0xd7af5f,
		0xd7afd7,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
