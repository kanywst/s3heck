package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GetX509Information(inputCertificate string, issuer bool, subject bool) (d string) {
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
	for _, c := range chain {
		cert, _ := x509.ParseCertificate(c)
		if issuer {
			d += fmt.Sprintf("Issuer CN: %s\n", cert.Issuer.CommonName)
		}
		if subject {
			d += fmt.Sprintf("Subject CN: %s\n", cert.Subject.CommonName)
		}
		d += "\n"
	}
	return
}
