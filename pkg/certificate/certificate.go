package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kanywst/s3heck/pkg/color"
)

func GetX509Information(inputCertificate string, issuer bool, subject bool, validity bool, dns bool) (d string) {
	var (
		chain     [][]byte
		certBlock *pem.Block
	)

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
			d += color.ColorStr("Issuer", cert.Issuer.CommonName, color.Colors[i%len(color.Colors)+1])
		}
		if subject {
			d += color.ColorStr("Subject", cert.Subject.CommonName, color.Colors[i%len(color.Colors)])
		}
		if validity {
			d += color.ColorStr("Validity", "", "")
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
