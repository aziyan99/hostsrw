package main

import (
	"fmt"
	"os"

	"github.com/aziyan99/hostsrw/pkg/helper"
	"github.com/aziyan99/hostsrw/pkg/hostsrw"
)

// TODO: Save conf to *.ini file(?)
// TODO: Add debug flag for -verbose output
// TODO: Add unit tests
func main() {
	args := os.Args

	if len(args) < 2 {
		helper.Help()
		os.Exit(1)
	}

	switch args[1] {
	case "all":
		hosts, err := hostsrw.All()
		if err != nil {
			helper.Check(err)
		}

		for i := 0; i < len(hosts); i++ {
			fmt.Println(hosts[i])
		}

	case "exists":
		// TODO: Also accept IP
		hosts, err := hostsrw.Exists(args[2])
		if err != nil {
			helper.Check(err)
		}

		for i := 0; i < len(hosts); i++ {
			fmt.Println(hosts[i])
		}
	case "add":
		if err := hostsrw.Add(args[2]); err != nil {
			helper.Check(err)
		}
	case "rm":
		if err := hostsrw.Remove(args[2]); err != nil {
			helper.Check(err)
		}
	default:
		helper.Help()
	}

	os.Exit(0)
}
