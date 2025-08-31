package hostsrw

import (
	"os"
	"strings"

	"github.com/aziyan99/hostsrw/v2/pkg/helper"
)

func All() ([]string, error) {
	hostsBuf, err := os.ReadFile(helper.HOSTS_PATH)
	if err != nil {
		return []string{}, err
	}

	hosts := strings.Split(string(hostsBuf), helper.NEW_LINE_FLAG)

	var allHosts []string
	for i := 0; i < len(hosts); i++ {
		allHosts = append(allHosts, hosts[i])
	}

	if len(allHosts) == 0 {
		return []string{}, nil
	}

	return allHosts, nil
}
