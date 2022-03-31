package luna

import (
	"sort"

	"github.com/taigrr/go-colorpalettes/simplecolor"
)

var colors = simplecolor.SimplePalette{0x002b2b,
	0x003f3f,
	0x005e5e,
	0x2aa198,
	0x2e8b57,
	0x4e4e4e,
	0x450000,
	0x789f00,
	0x780000,
	0x973d45,
	0xe20000,
	0xff8036,
	0xffff9a,
	0xffffff,
}

func GetPalette() (sp simplecolor.SimplePalette) {
	for _, x := range colors {
		sp = append(sp, x)
	}
	sort.Sort(sp)
	return
}
