package main

import (
	"fmt"
	"strings"

	"github.com/taigrr/simplecolorpalettes/palettes/vim/luna"
)

func main() {
	colors := luna.GetPalette()
	colors = colors.ToExtendedAnsi()
	var open []string
	var close []string
	for _, c := range colors {
		red, green, blue, _ := c.RGBA()
		if (float32(red)*0.299 + float32(green)*0.587 + float32(blue)*0.114) > 150.0 {
			open = append(open, fmt.Sprintf("\u001B[38;2;%d;%d;%dm", 0, 0, 0))
		} else {
			open = append(open, fmt.Sprintf("\u001B[38;2;%d;%d;%dm", 255, 255, 255))
		}
		close = append(close, "\u001B[39m")

		open = append(open, fmt.Sprintf("\u001B[48;2;%d;%d;%dm", red, green, blue))
		close = append(close, "\u001B[49m")

		//		if m.isUnderlined {
		// open = append(open, "\u001B[4m")
		// close = append(close, "\u001B[24m")
		//			}

		//			if m.isDim {
		//				open = append(open, "\u001B[2m")
		//				close = append(close, "\u001B[22m")
		//			}
		//
		//			if m.isItalic {
		//				open = append(open, "\u001B[3m")
		//				close = append(close, "\u001B[23m")
		//			}

		//			if m.isBold {
		//		open = append(open, "\u001B[1m")
		//		close = append(close, "\u001B[21m")
		//			}
		var b strings.Builder
		fmt.Fprint(&b, strings.Join(open, ""))
		fmt.Fprint(&b, "adasdasd")
		fmt.Fprint(&b, strings.Join(close, ""))

		fmt.Print(b.String() + " ")
	}
	fmt.Println("")
}
