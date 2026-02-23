package charmtone

import (
	"sort"
	"sync"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	once   sync.Once
	colors = simplecolor.SimplePalette{
		0xBF976F, // cumin
		0xFF985A, // tang
		0xFFB587, // yam
		0xD36C64, // paprika
		0xFF6E63, // bengal
		0xFF937D, // uni
		0xEB4268, // sriracha
		0xFF577D, // coral
		0xFF7F90, // salmon
		0xE23080, // chili
		0xFF388B, // cherry
		0xFF6DAA, // tuna
		0xE940B0, // macaron
		0xFF4FBF, // pony
		0xFF79D0, // cheeky
		0xF947E3, // flamingo
		0xFF60FF, // dolly
		0xFF84FF, // blush
		0xC337E0, // urchin
		0xEB5DFF, // crystal
		0xF379FF, // lilac
		0x9C35E1, // prince
		0xC259FF, // violet
		0xD46EFF, // mauve
		0x7134DD, // grape
		0x9953FF, // plum
		0xAD6EFF, // orchid
		0x4A30D9, // jelly
		0x6B50FF, // charple
		0x8B75FF, // hazy
		0x3331B2, // ox
		0x4949FF, // sapphire
		0x7272FF, // guppy
		0x2B55B3, // oceania
		0x4776FF, // thunder
		0x719AFC, // anchovy
		0x007AB8, // damson
		0x00A4FF, // malibu
		0x4FBEFE, // sardine
		0x10B1AE, // zinc
		0x0ADCD9, // turtle
		0x5CDFEA, // lichen
		0x12C78F, // guac
		0x00FFB2, // julep
		0x68FFD6, // bok
		0xF5EF34, // mustard
		0xE8FF27, // citron
		0xE8FE96, // zest
		0x121117, // caviar
		0x201F26, // pepper
		0x3A3943, // charcoal
		0x4D4C57, // iron
		0x605F6B, // oyster
		0x858392, // squid
		0xBFBCC8, // smoke
		0xDFDBDD, // ash
		0xF1EFEF, // salt
		0xFFFAF1, // butter
	}
)

func GetPalette() (sp simplecolor.SimplePalette) {
	once.Do(func() {
		sort.Sort(colors)
	})
	return colors
}
