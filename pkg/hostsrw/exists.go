package hostsrw

import (
	"os"
	"strings"
)

func Exists(entry string, hostPath string, newLineFlag string) ([]string, error) {
	hostsBuf, err := os.ReadFile(hostPath)
	if err != nil {
		return []string{}, err
	}

	hosts := strings.Split(string(hostsBuf), newLineFlag)

	var allHosts []string
	for i := 0; i < len(hosts); i++ {
		if strings.Contains(hosts[i], entry) {
			allHosts = append(allHosts, hosts[i])
		}
	}

	if len(allHosts) == 0 {
		return []string{}, nil
	}

	return allHosts, nil

}
