package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kanywst/s3heck/pkg/certinfo"
)

func GetX509Information(inputCertificate string, issuer bool, subject bool, validity bool, dns bool, short bool) (d string) {
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
		s := fmt.Sprintf("Certificate: SerialNumber is %x\n", cert.SerialNumber)
		d += s
		if short {
			sanDns := strings.Join(cert.DNSNames, ", ")
			d += certinfo.PrintFgColor("Issuer", cert.Issuer.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)+1])
			if len(sanDns) != 0 {
				d += certinfo.PrintFgColor("Subject", cert.Subject.CommonName+" "+strings.Join(cert.DNSNames, ", "), certinfo.FgColors[i%len(certinfo.FgColors)])
			} else {
				d += certinfo.PrintFgColor("Subject", cert.Subject.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)])
			}
			d += certinfo.PrintValidity(cert.NotBefore, cert.NotAfter)
		} else {
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
		for i := 0; i < len(s)-1; i++ {
			d += "-"
		}
		d += "\n"
	}
	return
}
