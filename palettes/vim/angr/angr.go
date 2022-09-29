package angr

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f87,
		0x303030,
		0x3a3a3a,
		0x444444,
		0x87afd7,
		0xafafd7,
		0xb2b2b2,
		0xd78787,
		0xd7afd7,
		0xffaf87,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
