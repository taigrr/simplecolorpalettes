package solarized

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
		0x268bd2,
		0x2aa198,
		0x586e75,
		0x657b83,
		0x6c71c4,
		0x839496,
		0x859900,
		0x93a1a1,
		0xb58900,
		0xcb4b16,
		0xd33682,
		0xdc322f,
		0xeee8d5,
		0xfdf6e3,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
