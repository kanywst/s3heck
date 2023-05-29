package certinfo

import (
	"fmt"
	"time"

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

func PrintFgColor(k string, v string, c string) (ss string) {
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

func PrintValidity(notBefore time.Time, notAfter time.Time) (ss string) {
	ss += fmt.Sprintf("%2sValidity\n", "")
	ss += fmt.Sprintf("%4sfrom : %2s\n", "", notBefore)
	ss += fmt.Sprintf("%4sto   : %2s\n", "", notAfter)
	return
}

func PrintX509Extensions(san string) (ss string) {
	ss += fmt.Sprintf("%2sX509v3 extensions\n", "")
	ss += fmt.Sprintf("%4sX509v3 Subject Alternative Name: %s\n", "", san)
	return
}
