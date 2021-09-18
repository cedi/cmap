package cmd

import (
	"fmt"

	kout "github.com/cedi/kkpctl/pkg/output"
	"github.com/spf13/cobra"
)

var versionCMD = &cobra.Command{
	Use:     "version",
	Short:   "Shows version information",
	Example: "cmap version",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		render := make([]kout.VersionRender, 0)

		// Prepare kkpctl binary version
		render = append(render, kout.VersionRender{
			Component: "cmap",
			Version:   Version,
			Date:      Date,
			Commit:    Commit,
			BuiltBy:   BuiltBy,
		})

		parsed, err := kout.ParseOutput(render, outputType, kout.Name)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Print(parsed)
	},
}

func init() {
	rootCmd.AddCommand(versionCMD)
}
