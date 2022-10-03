package kolor

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005154,
		0x242322,
		0x4a4a4a,
		0x4f3598,
		0x75d7d8,
		0x7eaefd,
		0x875faf,
		0xb2b2b2,
		0xd96e8a,
		0xdbc570,
		0xe2e2e2,
		0xe6987a,
		0xff5fd7,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
