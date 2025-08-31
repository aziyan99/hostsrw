package hostsrw

import (
	"fmt"
	"os"
	"strings"

	"github.com/aziyan99/hostsrw/pkg/helper"
)

func Exists(entry string, hostPath string, newLineFlag string) {
	hostsBuf, err := os.ReadFile(hostPath)
	helper.Check(err)

	hosts := strings.Split(string(hostsBuf), newLineFlag)

	for i := 0; i < len(hosts); i++ {
		if strings.Contains(hosts[i], entry) {
			fmt.Printf("%s\n", hosts[i])
		}
	}
}
