package cmd

import (
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a specific ressource",
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
