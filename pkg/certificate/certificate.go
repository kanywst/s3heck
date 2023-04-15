package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kanywst/s3heck/pkg/certinfo"
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
		d += fmt.Sprintf("Certificate [%d]\n", i)
		if issuer {
			d += certinfo.PrintFgColor("Issuer", cert.Issuer.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)+1])
		}
		if subject {
			d += certinfo.PrintFgColor("Subject", cert.Subject.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)])
		}
		if validity {
			d += certinfo.PrintValidity(cert.NotBefore, cert.NotAfter)
		}
		if dns {
			d += certinfo.PrintX509Extensions(strings.Join(cert.DNSNames, ", "))
		}
	}
	return
}
