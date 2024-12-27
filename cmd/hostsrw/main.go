package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// TODO: Save conf to *.ini file(?)
// TODO: Add debug flag for -verbose output
const (
	HostsRWVersion = "1.0.0"

	HOSTS_PATH    = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	NEW_LINE_FLAG = "\n"
	HOSTS_FLAG    = "yawamp"
)

// TODO: Add unit tests
func main() {
	args := os.Args

	if !isTerminal() {
		os.Exit(1)
	}

	if len(args) < 3 {
		help()
		return
	}

	switch args[1] {
	case "exists":
		// TODO: Also accept IP
		exists(args[2])
	case "add":
		add(args[2])
	case "rm":
		remove(args[2])
	default:
		help()
	}

}

func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}

	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func help() {
	fmt.Println("\nhostsrw - read/write windows hosts file [version  " + HostsRWVersion + "]")
	fmt.Println("")
	fmt.Println("\nUsage:")
	fmt.Println(" ")
	fmt.Println("  hostsrw add [entry]        : Add a new entry.")
	fmt.Println("  hostsrw rm  [entry]        : Remove an existng entry.")
	fmt.Println("  hostsrw exists [entry]     : Determine if entry is exists.")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\nExample:")
	fmt.Println(" ")
	fmt.Println("  hostsrw add foo.test       : Add '127.0.0.1 foo.test' to hosts")
	fmt.Println("  hostsrw rm foo.test        : Remove '127.0.0.1 foo.test' from hosts")
	fmt.Println("  hostsrw exists foo.test    : Will search foo 'foo.test'")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func runMeElevated() {
	// https://gist.github.com/jerblack/d0eb182cc5a1c1d92d92a4c4fcc416c6
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func exists(entry string) {
	hostsBuf, err := os.ReadFile(HOSTS_PATH)
	check(err)

	hosts := strings.Split(string(hostsBuf), NEW_LINE_FLAG)

	for i := 0; i < len(hosts); i++ {
		if strings.Contains(hosts[i], entry) {
			fmt.Printf("%s\n", hosts[i])
		}
	}
}

func remove(entry string) {
	hostsBuf, err := os.ReadFile(HOSTS_PATH)
	check(err)

	hosts := strings.Split(string(hostsBuf), NEW_LINE_FLAG)

	var newHosts []string
	for i := 0; i < len(hosts); i++ {
		if !strings.Contains(hosts[i], entry) {
			newHosts = append(newHosts, hosts[i])
		}
	}

	// https://golang.org/pkg/os/#Chmod
	// On Windows, the mode must be non-zero but otherwise
	// only the 0200 bit (owner writable) of mode is used;
	// it controls whether the file's read-only attribute is set or cleared.
	// The other bits are currently unused.
	// Use mode 0400 for a read-only file and 0600 for a readable+writable file.

	f, _ := os.OpenFile(HOSTS_PATH, os.O_RDWR|os.O_TRUNC, 0600)

	defer f.Close()

	_, err = f.WriteString(strings.Join(newHosts, "\n"))
	if err != nil {
		if !amAdmin() {
			// TODO: Only ask permission when neccessary
			runMeElevated()
		}
	}

	f.Sync()
}

func add(entry string) {
	hostsBuf, err := os.ReadFile(HOSTS_PATH)
	check(err)

	hosts := strings.Split(string(hostsBuf), NEW_LINE_FLAG)

	newHost := "127.0.0.1\t\t" + entry + "\t\t\t\t#yawamp magic!\n"

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

	f, _ := os.OpenFile(HOSTS_PATH, os.O_RDWR|os.O_TRUNC, 0600)

	defer f.Close()

	_, err = f.WriteString(strings.Join(newHosts, "\n"))
	if err != nil {
		// TODO: Only ask permission when neccessary
		if !amAdmin() {
			runMeElevated()
		}
	}

	f.Sync()
}
