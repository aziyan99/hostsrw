package main

import (
	"os"

	"github.com/aziyan99/hostsrw/pkg/helper"
	"github.com/aziyan99/hostsrw/pkg/hostsrw"
)

// TODO: Save conf to *.ini file(?)
// TODO: Add debug flag for -verbose output
const (
	HostsRWVersion = "1.1.0"

	HOSTS_PATH    = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	NEW_LINE_FLAG = "\n"
)

// TODO: Add unit tests
func main() {
	args := os.Args

	if len(args) < 2 {
		helper.Help(HostsRWVersion)
		os.Exit(1)
	}

	switch args[1] {
	case "all":
		hostsrw.All(HOSTS_PATH, NEW_LINE_FLAG)
	case "exists":
		// TODO: Also accept IP
		hostsrw.Exists(args[2], HOSTS_PATH, NEW_LINE_FLAG)
	case "add":
		hostsrw.Add(args[2], HOSTS_PATH, NEW_LINE_FLAG)
	case "rm":
		hostsrw.Remove(args[2], HOSTS_PATH, NEW_LINE_FLAG)
	default:
		helper.Help(HostsRWVersion)
	}

	os.Exit(0)
}
