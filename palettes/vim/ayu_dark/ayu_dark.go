package ayu_dark

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0a0e14,
		0x304357,
		0x39bae6,
		0x3d424d,
		0xb3b1ad,
		0xc2d94c,
		0xff3333,
		0xff8f40,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
