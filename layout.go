package main

import tl "github.com/JoelOtter/termloop"

type xy struct {
	x, y int
}

type block struct {
	w, h       int
	marginTop  int
	marginLeft int
}

type Layout struct {
	square   block
	board    block
	piece    block
	ranks    block
	files    block
	notation block
	input    block
}

type Palette struct {
	dark, light, pieces, notations, current tl.Attr
	valid, invalid                          tl.Attr
	input                                   tl.Attr
}

var (
	layout = Layout{
		square: block{
			w: 7, h: 3,
		},
		board: block{
			w:          8,
			h:          8,
			marginLeft: 2,
		},
		piece: block{
			marginTop:  1,
			marginLeft: 2,
		},
		ranks: block{
			marginTop: 1,
		},
		files: block{
			marginLeft: 3,
		},
		notation: block{
			marginLeft: 3,
		},
		input: block{
			marginTop: 2,
		},
	}

	palette = Palette{
		dark:      tl.ColorCyan,
		light:     tl.ColorWhite,
		pieces:    tl.ColorBlack,
		notations: tl.ColorCyan,
		current:   tl.ColorDefault,
		valid:     tl.ColorYellow,
		invalid:   tl.ColorRed,
		input:     tl.ColorWhite,
	}
)

func (s Palette) alternate() tl.Attr {
	if palette.current == palette.light {
		palette.current = palette.dark
		return palette.current
	}
	palette.current = palette.light
	return palette.current
}
