package main

import (
	"fmt"
	"regexp"
)

// Originally from https://github.com/mitchellh/cli/blob/master/ui_colored.go

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

func (c *Color) Colorize(msg string) string {
	if c.Code == ColorDefault.Code {
		return msg
	}

	attr := 0
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", attr, c.Code, msg)
}
