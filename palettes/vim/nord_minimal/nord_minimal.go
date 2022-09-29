package nord_minimal

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x2e3440,
		0x3b4252,
		0x434c5e,
		0x4c566a,
		0x5e81ac,
		0x81a1c1,
		0x88c0d0,
		0x8fbcbb,
		0xa3be8c,
		0xb48ead,
		0xbf616a,
		0xd08770,
		0xd8dee9,
		0xe5e9f0,
		0xebcb8b,
		0xeceff4,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
