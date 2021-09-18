# cmap

[![GitHub license](https://img.shields.io/github/license/cedi/cmap.svg)](https://github.com/cedi/cmap/blob/main/LICENSE)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/cedi/cmap.svg)](https://github.com/cedi/cmap)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/cedi/cmap)
[![GoReportCard example](https://goreportcard.com/badge/github.com/cedi/cmap)](https://goreportcard.com/report/github.com/cedi/cmap)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/cedi/cmap.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/cedi/cmap/alerts/)
[![workflow status](https://github.com/cedi/cmap/actions/workflows/go.yml/badge.svg)](https://github.com/cedi/cmap/actions)

cmap is a poor-mans nmap version, that utilizes nmap under the hood.

I'm just really fucking lazy with learning the nmap flags properly, and 99% of the time, I just wanna know which hosts are up, and if port 22 is open for SSH. So yeah, call me stupid, but I actually use this tool...

## Usage

```bash
$ cmap scan network 192.168.0.0/24
  NAME (5)              IP              SSH PORT OPEN
 --------------------- --------------- ---------------
  router                192.168.0.1     Yes
  ava                   192.168.0.141   No
  clusterpi-master      192.168.0.156   Yes
  clusterpi-worker-1    192.168.0.111   Yes
  clusterpi-worker-2    192.168.0.103   Yes
```

## Pull requests

I warmly welcome pull requests. Feel free to contribute whatever you feel like (if you feel like it at all 😂 it's a kinda stupid tool anyway)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
