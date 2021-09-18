package cmd

import (
	"fmt"

	"github.com/cedi/cmap/pkg/scan"
	kout "github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var scanHostCmd = &cobra.Command{
	Use:     "host ip",
	Short:   "Scans a network",
	Example: "cmap scan host 192.168.0.123",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		result, err := scan.HostSSH(args[0])
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
	scanCmd.AddCommand(scanHostCmd)
}
