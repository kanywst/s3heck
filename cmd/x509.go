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
}

var (
	o = &Options{}
)

var x509Cmd = &cobra.Command{
	Use:   "x509",
	Short: "display necessary information about x509 certificates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(certificate.GetX509Information(o.certificateFileName, o.issuerCommonName, o.subjectCommonName, o.validity, o.dns))
	},
}

func init() {
	rootCmd.AddCommand(x509Cmd)

	x509Cmd.Flags().StringVarP(&o.certificateFileName, "pem", "p", "", "specify certificate input pem")
	x509Cmd.Flags().BoolVarP(&o.issuerCommonName, "issuer", "i", false, "display issuer")
	x509Cmd.Flags().BoolVarP(&o.subjectCommonName, "subject", "s", false, "display subject")
	x509Cmd.Flags().BoolVarP(&o.validity, "validity", "v", false, "display validity")
	x509Cmd.Flags().BoolVarP(&o.dns, "dns", "d", false, "display subject alternative name")

}
