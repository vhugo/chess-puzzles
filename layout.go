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

	colorScheme = map[string]Palette{
		"standard": Palette{
			dark:      tl.RgbTo256Color(73, 163, 179),
			light:     tl.ColorWhite,
			pieces:    tl.RgbTo256Color(0, 0, 0),
			notations: tl.RgbTo256Color(73, 163, 179),
			valid:     tl.RgbTo256Color(200, 200, 0),
			invalid:   tl.RgbTo256Color(200, 0, 0),
			input:     tl.ColorWhite,
		},
		"brown": Palette{
			dark:      tl.RgbTo256Color(175, 138, 105),
			light:     tl.RgbTo256Color(218, 218, 185),
			pieces:    tl.RgbTo256Color(0, 0, 0),
			notations: tl.RgbTo256Color(175, 138, 105),
			valid:     tl.RgbTo256Color(200, 200, 0),
			invalid:   tl.RgbTo256Color(200, 0, 0),
			input:     tl.ColorWhite,
		},
		"green": Palette{
			dark:      tl.RgbTo256Color(144, 173, 105),
			light:     tl.RgbTo256Color(196, 196, 196),
			pieces:    tl.RgbTo256Color(0, 0, 0),
			notations: tl.RgbTo256Color(144, 173, 105),
			valid:     tl.RgbTo256Color(200, 200, 0),
			invalid:   tl.RgbTo256Color(200, 0, 0),
			input:     tl.ColorWhite,
		},
		"dark": Palette{
			dark:      tl.RgbTo256Color(50, 50, 50),
			light:     tl.RgbTo256Color(150, 150, 150),
			pieces:    tl.RgbTo256Color(255, 255, 255),
			notations: tl.RgbTo256Color(150, 150, 150),
			valid:     tl.RgbTo256Color(180, 120, 0),
			invalid:   tl.RgbTo256Color(100, 0, 0),
			input:     tl.RgbTo256Color(255, 255, 255),
		},
	}

	palette = colorScheme["dark"]
)

func (s Palette) alternate() tl.Attr {
	if palette.current == palette.light {
		palette.current = palette.dark
		return palette.current
	}
	palette.current = palette.light
	return palette.current
}
