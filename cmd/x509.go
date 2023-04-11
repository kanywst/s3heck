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
}

var (
	o = &Options{}
)

// x509Cmd represents the x509 command
var x509Cmd = &cobra.Command{
	Use:   "x509",
	Short: "display x509 certificate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(certificate.GetX509Information(o.certificateFileName, o.issuerCommonName, o.subjectCommonName, o.validity))
	},
}

func init() {
	rootCmd.AddCommand(x509Cmd)

	x509Cmd.Flags().StringVarP(&o.certificateFileName, "file", "f", "", "Certificate input file")
	x509Cmd.Flags().BoolVarP(&o.issuerCommonName, "issuer", "i", false, "Display issuer")
	x509Cmd.Flags().BoolVarP(&o.subjectCommonName, "subject", "s", false, "Display subject")
	x509Cmd.Flags().BoolVarP(&o.validity, "validity", "v", false, "Display validity")
}
