package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aziyan99/hostsrw/v2/pkg/elevated"
	"github.com/aziyan99/hostsrw/v2/pkg/helper"
	"github.com/aziyan99/hostsrw/v2/pkg/hostsrw"
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
		helper.Check(err)

		for i := 0; i < len(hosts); i++ {
			fmt.Println(hosts[i])
		}

	case "exists":
		// TODO: Also accept IP
		hosts, err := hostsrw.Exists(args[2])
		helper.Check(err)

		for i := 0; i < len(hosts); i++ {
			fmt.Println(hosts[i])
		}
	case "add":
		if !elevated.AmAdmin() {
			elevated.RunMeElevated()
		}

		err := hostsrw.Add(args[2])
		if err != nil {
			log.Fatalf("unable to write '%s' into hosts file. Error: %v\n", args[2], err)
		}
	case "rm":
		if !elevated.AmAdmin() {
			elevated.RunMeElevated()
		}

		err := hostsrw.Remove(args[2])
		if err != nil {
			log.Fatalf("unable to remove '%s' from hosts file. Error: %v\n", args[2], err)
		}
	default:
		helper.Help()
	}

	os.Exit(0)
}
