package violet

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x009966,
		0x3a3a3a,
		0x4e4e4e,
		0x5f0000,
		0x875faf,
		0xbcbcbc,
		0xc6c6c6,
		0xcacfd2,
		0xce537a,
		0xd75fd7,
		0xff5faf,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
