package badwolf

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x000000,
		0x005fff,
		0x0a9dff,
		0x141413,
		0x242321,
		0x45413b,
		0x666462,
		0x8cffba,
		0xaeee00,
		0xb88853,
		0xc7915b,
		0xf4cf86,
		0xfade3e,
		0xff2c4b,
		0xff9eb8,
		0xffa724,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
