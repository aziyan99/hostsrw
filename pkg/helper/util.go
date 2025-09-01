package helper

import "fmt"

var (
	HostsRWVersion = "2.3.2"

	HOSTS_PATH    = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	NEW_LINE_FLAG = "\n"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Help() {
	fmt.Println("\nhostsrw - read/write windows hosts file [version  " + HostsRWVersion + "]")
	fmt.Println("")
	fmt.Println("\nUsage:")
	fmt.Println(" ")
	fmt.Println("  hostsrw all                : List all lines in hosts file")
	fmt.Println("  hostsrw add [entry]        : Add a new entry.")
	fmt.Println("  hostsrw rm  [entry]        : Remove an existng entry.")
	fmt.Println("  hostsrw exists [entry]     : Determine if entry is exists.")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\nExample:")
	fmt.Println(" ")
	fmt.Println("  hostsrw all                : Will list all lines in hosts file")
	fmt.Println("  hostsrw add foo.test       : Add '127.0.0.1 foo.test' to hosts")
	fmt.Println("  hostsrw rm foo.test        : Remove '127.0.0.1 foo.test' from hosts")
	fmt.Println("  hostsrw exists foo.test    : Will search foo 'foo.test'")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
}
