package scan

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/cedi/cmap/pkg/output"
	"github.com/pkg/errors"
)

func Network(cidrs ...string) ([]output.HostShort, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -sn 192.168.0.0/24`, with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(cidrs...),
		nmap.WithContext(ctx),
		nmap.WithCustomArguments("-sn"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create nmap scanner")
	}

	scanResult, warnings, err := scanner.Run()
	if err != nil {
		return nil, errors.Wrap(err, "unable to run nmap scan")
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	result := make([]output.HostShort, 0)
	for _, host := range scanResult.Hosts {
		if host.Status.State != "up" {
			continue
		}

		host, err := HostSSH(host.Addresses[0].Addr)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to scan if Host has SSH Open")
		}

		result = append(result, host)
	}

	return result, nil
}

func HostSSH(host string) (output.HostShort, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -sn 192.168.0.0/24`, with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(host),
		nmap.WithContext(ctx),
		nmap.WithPorts("22"),
	)
	if err != nil {
		return output.HostShort{}, errors.Wrap(err, "unable to create nmap scanner")
	}

	scanResult, warnings, err := scanner.Run()
	if err != nil {
		return output.HostShort{}, errors.Wrap(err, "unable to run nmap scan")
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	if len(scanResult.Hosts) == 0 {
		return output.HostShort{}, fmt.Errorf("ScanHostSSH: no result returned for host %s", host)
	}

	hostObj := scanResult.Hosts[0]
	ips := make([]string, len(hostObj.Addresses))
	for idx, ip := range hostObj.Addresses {
		ips[idx] = ip.Addr
	}

	hostnames := make([]string, len(hostObj.Hostnames))
	for idx, hostname := range hostObj.Hostnames {
		hostnames[idx] = hostname.Name
	}

	if len(hostObj.Ports) == 0 {
		return output.HostShort{}, fmt.Errorf("ScanHostSSH: no result returned for port scan on host %s", host)
	}

	return output.HostShort{
		Name: strings.Join(hostnames, ", "),
		IP:   strings.Join(ips, ", "),
		SSH:  string(hostObj.Ports[0].Status()),
	}, nil
}
