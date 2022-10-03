package onedark

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x282c34,
		0x3e4452,
		0x61afef,
		0x98c379,
		0xabb2bf,
		0xc678dd,
		0xe06c75,
		0xe5c07b,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
