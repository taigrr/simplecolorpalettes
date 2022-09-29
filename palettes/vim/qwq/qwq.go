package qwq

import (
	"sort"
	"sync"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0x008492,
		0x04bec3,
		0x0e3b4f,
		0x6bad3f,
		0x8befc7,
		0x9ed47b,
		0xa28e79,
		0xb9a695,
		0xc1f9cd,
		0xddc6af,
		0xdde58e,
		0xf7846e,
		0xfefcf9,
		0xff003f,
		0xff5b6f,
		0xff5d4f,
		0xff9da5,
		0xffd3cb,
		0xffeee5,
		0xfff5d9,
		0xffffff,
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
