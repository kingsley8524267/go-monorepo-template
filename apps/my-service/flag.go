package my_service

import (
	"flag"
	"fmt"
	"os"
)

var (
	fVersion bool
)

func init() {
	flag.BoolVar(&fVersion, "v", false, "show version")
	flag.Parse()

	flagHandler()
}

func flagHandler() {
	if fVersion {
		fmt.Printf("%s: %s\n", name, version)
		os.Exit(0)
	}
}
