package sol

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x004b9a,
		0x007fff,
		0x09643f,
		0x343434,
		0x777777,
		0xa0a0a0,
		0xa3a3a3,
		0xb0b0b0,
		0xb3b3b3,
		0xc7c7c7,
		0xe33900,
		0xeeeeee,
		0xff2121,
		0xff3535,
		0xff6003,
		0xff6868,
		0xffdbc7,
		0xffff9a,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
