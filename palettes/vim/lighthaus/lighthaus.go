package lighthaus

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x00bfa4,
		0x090b26,
		0x21252d,
		0x50c16e,
		0xcccccc,
		0xd68eb2,
		0xe25600,
		0xed722e,
		0xfc2929,
		0xff4d00,
		0xff5050,
		0xfffade,
		0xffff00,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
