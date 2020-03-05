package main

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strings"
)

var colorByName = map[string]Color{
	"default": ColorDefault,
	"red":     ColorRed,
	"green":   ColorGreen,
	"yellow":  ColorYellow,
	"blue":    ColorBlue,
	"magenta": ColorMagenta,
	"cyan":    ColorCyan,
    "white":   ColorWhite,
    "grey":    ColorGrey,
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
				line = colorizer.Color.Colorize(line)
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
