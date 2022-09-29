package seagull

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x0099ff,
		0x00a5ab,
		0x0b141a,
		0x11ab00,
		0x1d252b,
		0x61707a,
		0x6d767d,
		0x787e82,
		0x808487,
		0x9854ff,
		0xbf8c00,
		0xe6eaed,
		0xff4053,
		0xff549b,
		0xff6200,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
