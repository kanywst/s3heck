package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var docsFS embed.FS

func SetDocsFS(fs embed.FS) {
	docsFS = fs
}

var learnCmd = &cobra.Command{
	Use:   "learn [topic]",
	Short: "Read documentation for included projects",
	Long:  "Learn provides access to the README files of the sub-projects (autocert, step-issuer, etc.) from embedded documentation.",
	Run: func(cmd *cobra.Command, args []string) {
		topics := map[string]string{
			"autocert":     "embedded_docs/autocert.md",
			"issuer":       "embedded_docs/step-issuer.md",
			"certificates": "embedded_docs/certificates.md",
			"cli":          "embedded_docs/cli.md",
			"mtls":         "embedded_docs/hello-mtls.md",
			"kms":          "embedded_docs/step-kms-plugin.md",
		}

		if len(args) == 0 {
			fmt.Println("Available topics:")
			for t := range topics {
				fmt.Printf(" - %s\n", t)
			}
			return
		}

		path, ok := topics[args[0]]
		if !ok {
			_, _ = fmt.Fprintln(os.Stderr, "Unknown topic. Try running 's3heck learn' to see available topics.")
			return
		}

		content, err := docsFS.ReadFile(path)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Could not read documentation for %s (path: %s): %v\n", args[0], path, err)
			return
		}

		fmt.Println(string(content))
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}
