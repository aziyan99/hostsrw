package helper

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Help(hostsRWVersion string) {
	fmt.Println("\nhostsrw - read/write windows hosts file [version  " + hostsRWVersion + "]")
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
