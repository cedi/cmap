package cmd

import (
	"fmt"

	"github.com/cedi/cmap/pkg/scan"
	kout "github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var scanNetworkCmd = &cobra.Command{
	Use:     "network cidr",
	Short:   "Scans a network",
	Example: "cmap scan network 192.168.0.0/24",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		result, err := scan.Network(args...)
		if err != nil {
			return errors.Wrap(err, "failed to scan network")
		}

		parsed, err := kout.ParseOutput(result, outputType, kout.Name)
		if err != nil {
			return errors.Wrap(err, "failed to parse network")
		}

		fmt.Print(parsed)

		return nil
	},
}

func init() {
	scanCmd.AddCommand(scanNetworkCmd)
}
