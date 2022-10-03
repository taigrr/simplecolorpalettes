package laederon

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x081c8c,
		0x1693a5,
		0x242321,
		0x2e2d2a,
		0x594512,
		0x90a680,
		0xab3e5b,
		0xab3e5d,
		0xef393d,
		0xf8f6f2,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
