package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
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

var colorByName = map[string]Color{
	"default": ColorDefault,
	"red":     ColorRed,
	"green":   ColorGreen,
	"yellow":  ColorYellow,
	"blue":    ColorBlue,
	"magenta": ColorMagenta,
	"cyan":    ColorCyan,
}

// colorex [COLOR:PATTERN, ...]
// | colorex 'yellow:status:4' 'red:status:5'
func main() {
	flag.Parse()

	colorizers := buildColorizers(flag.Args())

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		for _, colorizer := range colorizers {
			if colorizer.Pattern.MatchString(line) {
				line = colorize(line, colorizer.Color)
				break
			}
		}

		os.Stdout.WriteString(line + "\n")
	}
}

func buildColorizers(args []string) []Colorizer {
	colorizers := make([]Colorizer, len(args)-1)

	for _, arg := range args {
		parts := strings.SplitN(arg, ":", 2)
		name, pattern := parts[0], parts[1]

		color := colorByName[name]
		re := regexp.MustCompile(pattern)
		colorizer := Colorizer{re, color}

		colorizers = append(colorizers, colorizer)
	}

	return colorizers
}

func colorize(msg string, color Color) string {
	if color.Code == ColorDefault.Code {
		return msg
	}

	attr := 0
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", attr, color.Code, msg)
}
