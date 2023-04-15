package color

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	FgColors = map[int]string{
		0: "red",
		1: "yellow",
		2: "blue",
		3: "green",
		4: "cyan",
	}
)

func FgColorStr(k string, v string, c string) (ss string) {
	switch c {
	case "red":
		red := color.New(color.FgRed).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", red(k), red(v))
	case "yellow":
		yellow := color.New(color.FgYellow).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", yellow(k), yellow(v))
	case "blue":
		blue := color.New(color.FgBlue).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", blue(k), blue(v))
	case "green":
		green := color.New(color.FgGreen).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", green(k), green(v))
	case "cyan":
		cyan := color.New(color.FgCyan).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", cyan(k), cyan(v))
	default:
		ss = fmt.Sprintf("%2s%s: %s\n", "", k, v)

	}
	return
}
