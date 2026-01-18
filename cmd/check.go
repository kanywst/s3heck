package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	checkFile   string
	checkRemote string
	checkK8s    bool
)

var (
	errorColor   = color.New(color.FgRed)
	warningColor = color.New(color.FgYellow)
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Smart diagnostics for Certificates, TLS, and Kubernetes",
	Long: `Analyze and troubleshoot PKI issues automatically.

Solving OpenSSL pain:
- Detects expired/expiring certificates
- Verifies chain completeness and hostname matching
- Identifies weak algorithms

Solving K8s Fragmentation:
- Scans Autocert/Step-Issuer pods logs and events for common errors
- Checks if MutatingWebhooks are properly configured`,
	Run: func(cmd *cobra.Command, args []string) {
		if checkK8s {

			diagnoseKubernetes()
			return
		}

		if checkFile != "" {

			diagnoseLocalFile(checkFile)
			return
		}

		if checkRemote != "" {

			diagnoseRemoteEndpoint(checkRemote)
			return
		}

		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&checkFile, "file", "f", "", "Path to x509 certificate file")
	checkCmd.Flags().StringVarP(&checkRemote, "remote", "r", "", "Remote host:port to diagnose")
	checkCmd.Flags().BoolVarP(&checkK8s, "k8s", "k", false, "Deep diagnose of Kubernetes PKI components")
}

// -----------------------------------------------------------------------------
// OpenSSL Replacement Logic
// -----------------------------------------------------------------------------

func diagnoseLocalFile(path string) {
	fmt.Printf("🔍 Analysis: %s\n", path)
	data, err := os.ReadFile(path)
	if err != nil {
		_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Error reading file: %v\n", err)
		return
	}

	var certs []*x509.Certificate
	for len(data) > 0 {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			break
		}
		if block.Type != "CERTIFICATE" {
			continue
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Corrupt x509 data: %v\n", err)
			continue
		}
		certs = append(certs, cert)
	}

	if len(certs) == 0 {
		_, _ = errorColor.Fprintln(os.Stderr, "  ✘ Not a valid PEM file or no certificates found")
		return
	}

	fmt.Printf("\n--- Certificate Analysis (%d certs found) ---\n", len(certs))
	for i, cert := range certs {
		analyzeCertQuality(cert, "")
		if i < len(certs)-1 {
			fmt.Println("\n----------------------------------------")
		}
	}
}

func diagnoseRemoteEndpoint(target string) {
	host, port, err := net.SplitHostPort(target)
	if err != nil {
		// If it's a missing port error, we can assume the target is the host and default the port.
		if addrErr, ok := err.(*net.AddrError); ok && strings.Contains(addrErr.Err, "missing port") {
			host = target
			port = "443"
		} else {
			// For any other error, we cannot safely parse, so we should fail.
			_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Invalid remote address '%s': %v\n", target, err)
			return
		}
	}
	if port == "" {
		port = "443"
	}
	address := net.JoinHostPort(host, port)

	fmt.Printf("🔍 Probing: %s\n", address)

	// Explicitly ask for the host to check SNI behavior
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", address, &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: true, // We verify manually to give better error messages
	})
	if err != nil {
		_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Connection Refused: %v\n", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Error closing connection: %v\n", err)
		}
	}()

	state := conn.ConnectionState()
	certs := state.PeerCertificates

	if len(certs) == 0 {
		_, _ = errorColor.Fprintln(os.Stderr, "  ✘ No certificates presented by server.")
		return
	}

	leaf := certs[0]
	fmt.Println("\n--- 🛡️  Leaf Certificate Analysis ---")
	analyzeCertQuality(leaf, host)

	fmt.Println("\n--- 🔗 Chain Analysis ---")
	// Verify certificate chain against system trust store
	intermediatePool := x509.NewCertPool()
	if len(certs) > 1 {
		for _, cert := range certs[1:] {
			intermediatePool.AddCert(cert)
		}
	}
	rootPool, err := x509.SystemCertPool()
	if err != nil {
		_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Could not load system root CA pool: %v\n", err)
	}
	opts := x509.VerifyOptions{
		// DNSName is checked in analyzeCertQuality, but we include it here for a full verification.
		DNSName:       host,
		Intermediates: intermediatePool,
		Roots:         rootPool,
	}

	if _, err := leaf.Verify(opts); err != nil {
		// HostnameError is already reported, so we only report other chain-related errors.
		if _, ok := err.(x509.HostnameError); !ok {
			_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Chain is not trusted: %v\n", err)
		}
	} else {
		color.Green("  ✔ Chain is trusted by the system.")
	}

	if len(certs) == 1 {
		_, _ = warningColor.Fprintln(os.Stderr, "  ⚠️  Chain is short (only 1 cert). Intermediate CA might be missing.")
		_, _ = warningColor.Fprintln(os.Stderr, "      Some clients (Android, Java) will fail to connect.")
	} else {
		for i, c := range certs[1:] {
			fmt.Printf("  %d. %s\n", i+1, c.Subject.CommonName)
		}
		color.Green("  ✔ Chain length looks okay (%d certs)", len(certs))
	}

	// Handshake protocol check
	fmt.Println("\n--- 🤝 Protocol Analysis ---")
	switch state.Version {
	case tls.VersionTLS13:
		color.Green("  ✔ TLS 1.3 (Modern & Secure)")
	case tls.VersionTLS12:
		color.Green("  ✔ TLS 1.2 (Standard)")
	case tls.VersionTLS11:
		_, _ = errorColor.Fprintln(os.Stderr, "  ✘ Legacy TLS Version (TLS 1.1). Deprecated and insecure!")
	case tls.VersionTLS10:
		_, _ = errorColor.Fprintln(os.Stderr, "  ✘ Legacy TLS Version (TLS 1.0). Deprecated and insecure!")
	default:
		_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Unknown TLS Version (0x%x)\n", state.Version)
	}
}

func analyzeCertQuality(cert *x509.Certificate, expectedHost string) {
	// 1. Expiration
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
	if daysLeft < 0 {
		_, _ = errorColor.Fprintf(os.Stderr, "  ✘ EXPIRED %d days ago (%s)\n", -daysLeft, cert.NotAfter.Format("2006-01-02"))
	} else if daysLeft < 30 {
		_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Expires soon: %d days left (%s)\n", daysLeft, cert.NotAfter.Format("2006-01-02"))
	} else {
		color.Green("  ✔ Validity: OK (%d days left)", daysLeft)
	}

	// 2. Hostname Match
	if expectedHost != "" {
		if err := cert.VerifyHostname(expectedHost); err != nil {
			_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Hostname Mismatch: Certificate is for %v, not '%s'\n", cert.DNSNames, expectedHost)
		} else {
			color.Green("  ✔ Hostname matches CN/SANs")
		}
	}

	// 3. Algorithm Strength
	switch cert.SignatureAlgorithm {
	case x509.MD2WithRSA, x509.MD5WithRSA, x509.SHA1WithRSA, x509.DSAWithSHA1, x509.ECDSAWithSHA1:
		_, _ = errorColor.Fprintf(os.Stderr, "  ✘ Weak Algorithm: %s (Vulnerable!)\n", cert.SignatureAlgorithm.String())
	default:
		color.Cyan("  ℹ️  Algorithm: %s", cert.SignatureAlgorithm.String())
	}

	// 4. CA Check
	if cert.IsCA {
		_, _ = warningColor.Fprintln(os.Stderr, "  ⚠️  This is a CA certificate (not a leaf).")
	}

	fmt.Printf("  ℹ️  Subject: %s\n", cert.Subject)
	fmt.Printf("  ℹ️  Issuer:  %s\n", cert.Issuer)
}

// -----------------------------------------------------------------------------
// K8s Fragmentation Solver
// -----------------------------------------------------------------------------

func diagnoseKubernetes() {
	fmt.Println("🔍 Deep Scanning Kubernetes Cluster...")

	if _, err := exec.LookPath("kubectl"); err != nil {
		_, _ = errorColor.Fprintln(os.Stderr, "  ✘ kubectl not found. Cannot inspect cluster.")
		return
	}

	var hasErrors bool

	// 1. Check Autocert
	if !checkController("Autocert", "app=autocert") {
		hasErrors = true
	}

	// 2. Check Step Issuer
	if !checkController("Step Issuer", "app.kubernetes.io/name=step-issuer") {
		// Not necessarily an error if not using it, but good to know
		fmt.Println("  (Note: Step Issuer not found)")
	}

	// 3. Scan for recent certificate errors in events
	fmt.Println("Scanning cluster events for certificate issues...")
	eventsJSON, err := exec.Command("kubectl", "get", "events", "-A", "-o", "json").Output()
	if err != nil {
		_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Could not get cluster events: %v\n", err)
	} else {
		var eventList struct {
			Items []struct {
				Type    string `json:"type"`
				Reason  string `json:"reason"`
				Message string `json:"message"`
				Object  struct {
					Name string `json:"name"`
				} `json:"involvedObject"`
			} `json:"items"`
		}

		if err := json.Unmarshal(eventsJSON, &eventList); err != nil {
			_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Could not parse cluster events: %v\n", err)
		} else {
			foundIssues := 0
			for _, event := range eventList.Items {
				msg := strings.ToLower(event.Message)
				reason := strings.ToLower(event.Reason)
				if strings.Contains(msg, "cert-manager") || strings.Contains(msg, "step") || strings.Contains(msg, "autocert") ||
					strings.Contains(reason, "cert") || strings.Contains(reason, "tls") {
					if event.Type != "Normal" {
						_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  [%s] %s: %s (on %s)\n", event.Type, event.Reason, event.Message, event.Object.Name)
						foundIssues++
					}
				}
			}
			if foundIssues == 0 {
				color.Green("  ✔ No recent certificate-related errors found in events.")
			} else {
				hasErrors = true
				fmt.Println("  💡 Hint: Check 'kubectl logs' for the objects mentioned above.")
			}
		}
	}

	// 4. Check Webhook Configuration (Common failure point for Autocert)
	fmt.Print("Checking MutatingWebhooks (Autocert injection)... ")
	webhooksJSON, err := exec.Command("kubectl", "get", "mutatingwebhookconfigurations", "-o", "json").Output()
	if err != nil {
		_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Could not get mutatingwebhookconfigurations: %v\n", err)
		hasErrors = true
	} else {
		var webhookList struct {
			Items []struct {
				Metadata struct {
					Name string `json:"name"`
				} `json:"metadata"`
			} `json:"items"`
		}

		if err := json.Unmarshal(webhooksJSON, &webhookList); err != nil {
			_, _ = warningColor.Fprintf(os.Stderr, "  ⚠️  Could not parse webhook configurations: %v\n", err)
			hasErrors = true
		} else {
			found := false
			for _, wh := range webhookList.Items {
				if strings.Contains(wh.Metadata.Name, "autocert") {
					found = true
					break
				}
			}
			if found {
				color.Green("Configured")
			} else {
				_, _ = warningColor.Fprintln(os.Stderr, "Not found (Autocert injection might fail)")
				hasErrors = true
			}
		}
	}

	// 5. Summary
	fmt.Println("\n--- Diagnosis Summary ---")
	if hasErrors {
		_, _ = errorColor.Fprintln(os.Stderr, "✘ Issues detected in your PKI infrastructure.")
		fmt.Println("  Try 's3heck learn autocert' or 's3heck learn issuer' for troubleshooting guides.")
	} else {
		color.Green("✔ Cluster PKI components look healthy.")
	}
}

func checkController(name, labelSelector string) bool {
	fmt.Printf("Checking %s Controller... ", name)
	out, err := exec.Command("kubectl", "get", "pods", "-A", "-l", labelSelector, "-o", "jsonpath={.items[*].status.phase}").Output()
	if err == nil && len(strings.TrimSpace(string(out))) > 0 {
		allRunning := true
		phases := strings.Fields(string(out))
		for _, phase := range phases {
			if phase != "Running" {
				allRunning = false
				break
			}
		}
		if allRunning {
			color.Green("Running (%d pods)", len(phases))
			return true
		} else {
			_, _ = errorColor.Fprintf(os.Stderr, "Partial Failure (phases: %s)\n", strings.Join(phases, ", "))
			return false
		}
	} else {
		_, _ = errorColor.Fprintln(os.Stderr, "Not Running or Not Found")
		return false
	}
}
