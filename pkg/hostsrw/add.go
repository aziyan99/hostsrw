package hostsrw

import (
	"os"
	"strings"

	"github.com/aziyan99/hostsrw/pkg/elevated"
	"github.com/aziyan99/hostsrw/pkg/helper"
)

func Add(entry string, hostPath string, newLineFlag string) {
	hostsBuf, err := os.ReadFile(hostPath)
	helper.Check(err)

	hosts := strings.Split(string(hostsBuf), newLineFlag)

	newHost := "127.0.0.1\t\t" + entry + "\n"

	// TODO: Ignore whitespace or empty lines when appending hosts

	var newHosts []string
	isEntryExists := false
	for i := 0; i < len(hosts); i++ {
		if strings.Contains(hosts[i], entry) {
			isEntryExists = true
		}

		newHosts = append(newHosts, hosts[i])
	}

	if !isEntryExists {
		newHosts = append(newHosts, newHost)
	}

	// https://golang.org/pkg/os/#Chmod
	// On Windows, the mode must be non-zero but otherwise
	// only the 0200 bit (owner writable) of mode is used;
	// it controls whether the file's read-only attribute is set or cleared.
	// The other bits are currently unused.
	// Use mode 0400 for a read-only file and 0600 for a readable+writable file.

	f, _ := os.OpenFile(hostPath, os.O_RDWR|os.O_TRUNC, 0600)

	defer f.Close()

	_, err = f.WriteString(strings.Join(newHosts, "\n"))
	if err != nil {
		// TODO: Only ask permission when neccessary
		if !elevated.AmAdmin() {
			elevated.RunMeElevated()
		}
	}

	f.Sync()
}
