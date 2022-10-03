package jellybeans

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0d61ac,
		0x151515,
		0x262626,
		0x437019,
		0x4f5b66,
		0x5fb3b3,
		0x65737e,
		0x666666,
		0x870000,
		0xa7adba,
		0xab7967,
		0xc0c5ce,
		0xc594c5,
		0xcdd3de,
		0xd8dee9,
		0xf99157,
		0xfac863,
		0xffb964,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
