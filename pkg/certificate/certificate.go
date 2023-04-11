package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GetX509Information(inputCertificate string, issuer bool, subject bool, validity bool) (d string) {
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
			d += fmt.Sprintf("%2sIssuer: %s\n", "", cert.Issuer.CommonName)
		}
		if subject {
			d += fmt.Sprintf("%2sSubject: %s\n", "", cert.Subject.CommonName)
		}
		if validity {
			d += fmt.Sprintf("%2sValidity:\n", "")
			d += fmt.Sprintf("%4sNotBefore: %s\n", "", cert.NotBefore)
			d += fmt.Sprintf("%4sNotAfter: %s\n", "", cert.NotAfter)
		}
	}
	return
}
