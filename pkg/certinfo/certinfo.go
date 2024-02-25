package certinfo

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	FgColors = []color.Attribute{
		color.FgRed,
		color.FgYellow,
		color.FgBlue,
		color.FgGreen,
		color.FgCyan,
	}
)

func PrintFgColor(k, v string, c color.Attribute) string {
	colorFunc := color.New(c).SprintFunc()
	return fmt.Sprintf("%2s%s: %s\n", "", colorFunc(k), colorFunc(v))
}

func PrintValidity(notBefore, notAfter time.Time) string {
	return fmt.Sprintf("%2sValidity\n%4sfrom : %s\n%4sto   : %s\n", "", "", notBefore, "", notAfter)
}

func PrintX509Extensions(san string) string {
	return fmt.Sprintf("%2sX509v3 extensions\n%4sX509v3 Subject Alternative Name: %s\n", "", "", san)
}
