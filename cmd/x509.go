package cmd

import (
	"fmt"

	"github.com/kanywst/s3heck/pkg/certificate"
	"github.com/spf13/cobra"
)

type Options struct {
	certificateFileName string
	issuerCommonName    bool
	subjectCommonName   bool
	validity            bool
	dns                 bool
	short               bool
}

var o = &Options{}

var x509Cmd = &cobra.Command{
	Use:   "x509",
	Short: "Output necessary information about x509 certificates",
	RunE: func(cmd *cobra.Command, args []string) error {
		if o.certificateFileName == "" {
			return fmt.Errorf("certificate file name not provided")
		}
		fmt.Print(certificate.GetX509Information(o.certificateFileName, o.issuerCommonName, o.subjectCommonName, o.validity, o.dns, o.short))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(x509Cmd)

	x509Cmd.Flags().StringVarP(&o.certificateFileName, "cert", "c", "", "specify server certificate")
	x509Cmd.Flags().BoolVarP(&o.issuerCommonName, "issuer", "i", false, "Output issuer")
	x509Cmd.Flags().BoolVarP(&o.subjectCommonName, "subject", "s", false, "Output subject")
	x509Cmd.Flags().BoolVarP(&o.validity, "validity", "v", false, "Output validity")
	x509Cmd.Flags().BoolVarP(&o.dns, "dns", "d", false, "Output subject alternative name")
	x509Cmd.Flags().BoolVarP(&o.short, "short", "t", false, "Output minimum Output information about minimum x509 certificates")
}
