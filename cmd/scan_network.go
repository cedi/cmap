package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	nmap "github.com/Ullaakut/nmap/v2"
	"github.com/cedi/cmap/pkg/output"
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
		cidr := make([]string, len(args))
		copy(cidr, args)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		// Equivalent to `/usr/local/bin/nmap -p 22 cidr...`, with a 5 minute timeout.
		scanner, err := nmap.NewScanner(
			nmap.WithTargets(cidr...),
			nmap.WithPorts("22"),
			nmap.WithContext(ctx),
		)
		if err != nil {
			log.Fatalf("unable to create nmap scanner: %v", err)
		}

		scanResult, warnings, err := scanner.Run()
		if err != nil {
			log.Fatalf("unable to run nmap scan: %v", err)
		}

		if warnings != nil {
			log.Printf("Warnings: \n %v", warnings)
		}

		result := make([]output.HostShort, 0)
		for _, host := range scanResult.Hosts {
			if host.Status.State != "up" {
				continue
			}

			ips := make([]string, len(host.Addresses))
			for idx, ip := range host.Addresses {
				ips[idx] = ip.Addr
			}

			hostnames := make([]string, len(host.Hostnames))
			for idx, hostname := range host.Hostnames {
				hostnames[idx] = hostname.Name
			}

			sshOpen := false
			for _, port := range host.Ports {
				if port.ID != 22 {
					continue
				}

				sshOpen = port.State.State != "closed"
			}

			result = append(result, output.HostShort{
				Name: strings.Join(hostnames, ","),
				IP:   strings.Join(ips, ","),
				SSH:  sshOpen,
			})
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
