package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	timeout         string
	additionalPorts string
	extraNmapArgs   string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a specific ressource",
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "5m", "The Timeout for the nmap scan operation. Default 5m. Format time as GoLang time.Duration")
	scanCmd.PersistentFlags().StringVarP(&additionalPorts, "ports", "p", "", "A comma separated list of additional ports to scan")
	scanCmd.PersistentFlags().StringVar(&extraNmapArgs, "extra-nmap-args", "", "Number of paralell host scans")
}

func GetTimeout() time.Duration {
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		return 5 * time.Minute
	}

	return duration
}
