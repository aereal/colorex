package main

import (
	"regexp"
)

type Colorizer struct {
	Pattern *regexp.Regexp
	Color   Color
}

type Color struct {
	Code int
}

var (
	ColorDefault Color = Color{-1}
	ColorRed           = Color{31}
	ColorGreen         = Color{32}
	ColorYellow        = Color{33}
	ColorBlue          = Color{34}
	ColorMagenta       = Color{35}
	ColorCyan          = Color{36}
)
