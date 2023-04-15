package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

func certStr(k string, v string, c string) (ss string) {
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

func GetX509Information(inputCertificate string, issuer bool, subject bool, validity bool, dns bool) (d string) {
	var (
		chain     [][]byte
		certBlock *pem.Block
	)

	colors := map[int]string{
		0: "red",
		1: "blue",
		2: "yellow",
	}
	certBytes, _ := ioutil.ReadFile(inputCertificate)
	for {
		certBlock, certBytes = pem.Decode(certBytes)
		if certBlock == nil {
			break
		}
		if certBlock.Type == "CERTIFICATE" {
			chain = append(chain, certBlock.Bytes)
		}
	}
	for i, c := range chain {
		cert, _ := x509.ParseCertificate(c)
		d += fmt.Sprintf("Certificate [%d]\n", i+1)
		if issuer {
			d += certStr("Issuer", cert.Issuer.CommonName, colors[i%len(colors)+1])
		}
		if subject {
			d += certStr("Subject", cert.Subject.CommonName, colors[i%len(colors)])
		}
		if validity {
			d += certStr("Validity", "", "")
			d += fmt.Sprintf("%4sNotBefore: %s\n", "", cert.NotBefore)
			d += fmt.Sprintf("%4sNotAfter: %s\n", "", cert.NotAfter)
		}
		if dns {
			d += fmt.Sprintf("%2sX509v3 extensions\n", "")
			d += fmt.Sprintf("%4sX509v3 Subject Alternative Name: %s\n", "", strings.Join(cert.DNSNames, ", "))
		}
	}
	return
}
