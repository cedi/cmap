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

func Network(ctx context.Context, paralellScans int, additionalPorts string, extraNmapArgs string, cidrs ...string) ([]output.HostShort, error) {
	if ctx == nil {
		newCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		ctx = newCtx
	}

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

	if len(warnings) > 0 {
		log.Printf("Warnings: \n %v", warnings)
	}

	if err != nil {
		return nil, errors.Wrap(err, "unable to run nmap scan")
	}

	totalJobs := 0

	for _, host := range scanResult.Hosts {
		if host.Status.State != "up" {
			continue
		}

		totalJobs++
	}

	result := make([]output.HostShort, totalJobs)
	jobResult := make(chan output.HostShort, totalJobs)
	jobs := make(chan nmap.Address, totalJobs)

	for worker := 0; worker < paralellScans; worker++ {
		go hostWorker(worker, ctx, additionalPorts, extraNmapArgs, jobs, jobResult)
	}

	for _, host := range scanResult.Hosts {
		if host.Status.State != "up" {
			continue
		}

		jobs <- host.Addresses[0]
	}

	close(jobs)

	for job := 0; job < totalJobs; job++ {
		result[job] = <-jobResult
	}

	return result, nil
}

func hostWorker(id int, ctx context.Context, additionalPorts string, extraNmapArgs string, jobs <-chan nmap.Address, results chan<- output.HostShort) {
	for job := range jobs {
		//log.Printf("Debug: Worker %d, started scan-job %s\n", id, job.Addr)

		host, err := Host(ctx, additionalPorts, extraNmapArgs, job.Addr)
		if err != nil {
			host = output.HostShort{
				Name:       "Unknown",
				IP:         job.Addr,
				SSH:        "",
				OtherPorts: "",
				Error:      errors.Wrap(err, "Unable to scan Host").Error(),
			}
		}

		//log.Printf("Debug: Worker %d, finished scan-job %s\n", id, job.Addr)
		results <- host
	}
}

func Host(ctx context.Context, additionalPorts string, extraNmapArgs string, host string) (output.HostShort, error) {
	if ctx == nil {
		newCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		ctx = newCtx
	}

	scanPorts := strings.Split(additionalPorts, ",")
	if len(scanPorts) == 1 && scanPorts[0] == "" {
		scanPorts[0] = "22"
	} else {
		scanPorts = append(scanPorts, "22")
	}

	// Equivalent to `/usr/local/bin/nmap 192.168.0.123 -p`, with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(host),
		nmap.WithContext(ctx),
		nmap.WithPorts(scanPorts...),
		nmap.WithCustomArguments(extraNmapArgs),
	)
	if err != nil {
		return output.HostShort{}, errors.Wrap(err, "unable to create nmap scanner")
	}

	scanResult, warnings, err := scanner.Run()

	if len(warnings) > 0 {
		log.Printf("Warnings: \n %v", warnings)
	}

	if err != nil {
		return output.HostShort{}, errors.Wrap(err, "unable to run nmap scan")
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

	sshStatus := ""
	additionalPortStatus := make([]string, 0)

	for _, port := range hostObj.Ports {
		if port.ID == 22 {
			sshStatus = string(port.Status())
			continue
		}

		for _, probe := range scanPorts {
			if fmt.Sprintf("%d", port.ID) == probe {
				additionalPortStatus = append(additionalPortStatus,
					fmt.Sprintf("%s/%d/%s=%s", port.Service, port.ID, port.Protocol, string(port.Status())),
				)
			}
		}
	}

	return output.HostShort{
		Name:       strings.Join(hostnames, ", "),
		IP:         strings.Join(ips, ", "),
		SSH:        sshStatus,
		OtherPorts: strings.Join(additionalPortStatus, "\n"),
	}, nil
}
