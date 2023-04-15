package color

import (
	"fmt"

	"github.com/fatih/color"
)

Colors := map[int]string{
	0: "red",
	1: "blue",
	2: "yellow",
}

func ColorStr(k string, v string, c string) (ss string) {
	switch c {
	case "red":
		red := color.New(color.FgRed).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", red(k), red(v))
	case "blue":
		blue := color.New(color.FgBlue).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", blue(k), blue(v))
	case "yellow":
		yellow := color.New(color.FgYellow).SprintFunc()
		ss = fmt.Sprintf("%2s%s: %s\n", "", yellow(k), yellow(v))
	default:
		ss = fmt.Sprintf("%2s%s: %s\n", "", k, v)

	}
	return
}
