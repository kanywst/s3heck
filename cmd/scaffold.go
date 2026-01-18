package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold [type]",
	Short: "Generate configuration templates (autocert, issuer, mtls)",
	Long: `Scaffold generates starting configurations for various Smallstep components.
Available types:
- autocert: K8s deployment with autocert annotations
- issuer: StepIssuer and CertificateRequest YAMLs
- mtls: A basic Go mTLS client/server pair`,
	ValidArgs: []string{"autocert", "issuer", "mtls"},
	Args:      cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "autocert":
			fmt.Print(autocertTemplate)
		case "issuer":
			fmt.Print(stepIssuerTemplate)
		case "mtls":
			if err := writeFile("mtls-server.go", mtlsServerTemplate); err != nil {
				return err
			}
			if err := writeFile("mtls-client.go", mtlsClientTemplate); err != nil {
				return err
			}
			fmt.Println("✅ Created mtls-server.go and mtls-client.go")
		default:
			return fmt.Errorf("unknown scaffold type. Use: autocert, issuer, mtls")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}

func writeFile(name, content string) error {
	if err := os.WriteFile(name, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing %s: %w", name, err)
	}
	return nil
}

const autocertTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-mtls
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-mtls
  template:
    metadata:
      annotations:
        autocert.step.sm/name: hello-mtls.default.svc.cluster.local
      labels:
        app: hello-mtls
    spec:
      containers:
      - name: hello-mtls
        image: smallstep/hello-mtls-server-go:latest
`

const stepIssuerTemplate = `apiVersion: certmanager.step.sm/v1beta1
kind: StepIssuer
metadata:
  name: step-issuer
  namespace: default
spec:
  # The CA URL where step-certificates is running
  url: https://step-certificates.default.svc.cluster.local
  
  # The base64 encoded version of the CA root certificate.
  # Obtain with: kubectl get -o jsonpath="{.data['root_ca\.crt']}" configmaps/step-certificates-certs | step base64
  caBundle: <BASE64_CA_BUNDLE>
  
  provisioner:
    name: admin
    # The provisioner Key ID (KID).
    # Obtain with: kubectl get -o jsonpath="{.data['ca\.json']}" configmaps/step-certificates-config | jq -r .authority.provisioners[0].key.kid
    kid: <KID>
    passwordRef:
      name: step-certificates-provisioner-password
      key: password
`

const mtlsServerTemplate = `package main

/**
 * mTLS Server Boilerplate
 * 
 * Required files to run this:
 * - ca.crt:     The Root CA certificate
 * - server.crt: The server certificate (signed by the CA)
 * - server.key: The server private key
 * 
 * To generate test certificates using 'step':
 *   step certificate create --profile root-ca "Test Root CA" ca.crt ca.key
 *   step certificate create --profile leaf --ca ca.crt --ca-key ca.key "localhost" server.crt server.key
 */

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello mTLS!"))
		}),
	}

	log.Println("Listening on :8443")
	log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}
`

const mtlsClientTemplate = `package main

/**
 * mTLS Client Boilerplate
 * 
 * Required files to run this:
 * - ca.crt:     The Root CA certificate
 * - client.crt: The client certificate (signed by the CA)
 * - client.key: The client private key
 * 
 * To generate test certificates using 'step':
 *   step certificate create --profile leaf --ca ca.crt --ca-key ca.key "client-user" client.crt client.key
 */

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS12,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}
	log.Printf("Response: %s", body)
}
`
