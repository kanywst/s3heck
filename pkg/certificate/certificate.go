package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/kanywst/s3heck/pkg/certinfo"
)

func GetX509Information(inputCertificate string, issuer, subject, validity, dns, short bool) string {
	file, err := os.Open(inputCertificate)
	if err != nil {
		return fmt.Sprintf("Error opening certificate file: %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Sprintf("Error getting file information: %v", err)
	}

	certBytes := make([]byte, stat.Size())
	_, err = file.Read(certBytes)
	if err != nil {
		return fmt.Sprintf("Error reading certificate file: %v", err)
	}

	var chain [][]byte
	for {
		certBlock, rest := pem.Decode(certBytes)
		if certBlock == nil {
			break
		}
		if certBlock.Type == "CERTIFICATE" {
			chain = append(chain, certBlock.Bytes)
		}
		certBytes = rest
	}

	var result strings.Builder
	for i, c := range chain {
		cert, err := x509.ParseCertificate(c)
		if err != nil {
			result.WriteString(fmt.Sprintf("Error parsing certificate: %v\n", err))
			continue
		}

		result.WriteString(fmt.Sprintf("Certificate: SerialNumber is %x\n", cert.SerialNumber))
		if short {
			issuerName := cert.Issuer.CommonName
			subjectName := cert.Subject.CommonName
			if len(cert.DNSNames) > 0 {
				subjectName += ", " + strings.Join(cert.DNSNames, ", ")
			}
			result.WriteString(certinfo.PrintFgColor("Issuer", issuerName, certinfo.FgColors[i%len(certinfo.FgColors)+1]))
			result.WriteString(certinfo.PrintFgColor("Subject", subjectName, certinfo.FgColors[i%len(certinfo.FgColors)]))
			result.WriteString(certinfo.PrintValidity(cert.NotBefore, cert.NotAfter))
		} else {
			if issuer {
				result.WriteString(certinfo.PrintFgColor("Issuer", cert.Issuer.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)+1]))
			}
			if subject {
				result.WriteString(certinfo.PrintFgColor("Subject", cert.Subject.CommonName, certinfo.FgColors[i%len(certinfo.FgColors)]))
			}
			if validity {
				result.WriteString(certinfo.PrintValidity(cert.NotBefore, cert.NotAfter))
			}
			if dns {
				result.WriteString(certinfo.PrintX509Extensions(strings.Join(cert.DNSNames, ", ")))
			}
		}
		result.WriteString(strings.Repeat("-", result.Len()-1) + "\n")
	}

	return result.String()
}
