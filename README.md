# cmap

[![GitHub license](https://img.shields.io/github/license/cedi/cmap.svg)](https://github.com/cedi/cmap/blob/main/LICENSE)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/cedi/cmap.svg)](https://github.com/cedi/cmap)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/cedi/cmap)
[![GoReportCard example](https://goreportcard.com/badge/github.com/cedi/cmap)](https://goreportcard.com/report/github.com/cedi/cmap)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/cedi/cmap.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/cedi/cmap/alerts/)
[![workflow status](https://github.com/cedi/cmap/actions/workflows/go.yml/badge.svg)](https://github.com/cedi/cmap/actions)

cmap is a poor-mans version of `nmap`. It utilizes nmap under the hood.

I'm just really fucking lazy learning the proper nmap syntax. Nost of the time, I just wanna know which hosts are up, and if port 22 is open for SSH. So yeah, call me stupid, but I actually use this tool almost on a daily base.
Sorry to all the nmap fans who feel offended by this creation of mine. But hey - at least it uses nmap under the hood, right? ;)

## Usage

```bash
$ cmap scan network 192.168.0.0/24 -p 6443
Starting scan of network [192.168.0.0/24] with Timeout 15m0s
  NAME (5)              IP              SSH        OTHER PORTS                    ERROR
 --------------------- --------------- ---------- ------------------------------ -------
  router                192.168.0.1     filtered   sun-sr-https/6443/tcp=closed
  ava                   192.168.0.140   closed     sun-sr-https/6443/tcp=closed
  clusterpi-master      192.168.0.156   open       sun-sr-https/6443/tcp=open
  clusterpi-worker-1    192.168.0.111   open       sun-sr-https/6443/tcp=closed
  clusterpi-worker-2    192.168.0.103   open       sun-sr-https/6443/tcp=closed
```

## Pull requests

I warmly welcome pull requests. Feel free to contribute whatever you feel like (if you feel like it at all 😂 it's a kinda stupid tool anyway)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

### Multi-Threading

It should be rather called "multi-go-routines". I utilize go-routines with a thread-pool design approach, so you can easily tailor the amount of threads used using the `--paralell` flag.
By default I spawn 10 go routines.

I don't actually multi-thread the network discovery, but only the host-detection part.
That means, if you wanna scan a `/12` network, it first calls `nmap -sn cidr/12` under the hood, which will take forever (or at least it will feel like it's taking forever). However, now that you got a couple hundred (or - imagine - a couple thousand) hosts reported as "online", the thread-pool is used to scan for open ports in paralell. Each thread is scanning one host at a time, using `nmap ipaddr -p 22`.
