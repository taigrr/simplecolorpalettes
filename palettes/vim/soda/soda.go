package soda

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x005f00,
		0x008700,
		0x00af00,
		0x5f0087,
		0x767676,
		0x875f87,
		0x875faf,
		0xdf0000,
		0xffaf5f,
		0xffd75f,
		0xffff5f,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
