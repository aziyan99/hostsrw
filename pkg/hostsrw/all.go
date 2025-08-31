package hostsrw

import (
	"fmt"
	"os"
	"strings"

	"github.com/aziyan99/hostsrw/pkg/helper"
)

func All(hostPath string, newLineFlag string) {
	hostsBuf, err := os.ReadFile(hostPath)
	helper.Check(err)

	hosts := strings.Split(string(hostsBuf), newLineFlag)

	for i := 0; i < len(hosts); i++ {
		fmt.Printf("%s\n", hosts[i])
	}
}
