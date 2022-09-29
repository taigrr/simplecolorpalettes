package solarized_flood

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x002b36,
		0x073642,
		0x1c1c1c,
		0x262626,
		0x268bd2,
		0x2aa198,
		0x303030,
		0x4e4e4e,
		0x586e75,
		0x5f00af,
		0x657b83,
		0x6c71c4,
		0x839496,
		0x859900,
		0x875faf,
		0x875fd7,
		0x93a1a1,
		0xb58900,
		0xcb4b16,
		0xd33682,
		0xd7d7ff,
		0xdc322f,
		0xeee8d5,
		0xfdf6e3,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
