package ouo

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005fff,
		0x00afff,
		0x1c1c1c,
		0x3a3a3a,
		0x4e4e4e,
		0x5faf00,
		0xaf0000,
		0xafaf87,
		0xd70000,
		0xd75f00,
		0xff0000,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
