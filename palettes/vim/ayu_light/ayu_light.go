package ayu_light

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x55b4d4,
		0x6c7680,
		0x86b300,
		0xf51818,
		0xfa8d3e,
		0xfafafa,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
