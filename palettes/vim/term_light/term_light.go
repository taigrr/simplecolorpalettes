package term_light

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x141413,
		0x242424,
		0x40403c,
		0x767676,
		0x7cb0e6,
		0x86cd74,
		0x888a85,
		0x94e42c,
		0xb5d3f3,
		0xcae682,
		0xdadada,
		0xdeded9,
		0xe55345,
		0xe5786d,
		0xf0f0f0,
		0xfade3e,
		0xfde76e,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
