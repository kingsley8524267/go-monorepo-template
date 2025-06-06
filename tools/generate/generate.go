package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./generate <app-name>")
		os.Exit(1)
	}

	appName := os.Args[1]

	var err error
	if err = genApp(appName); err != nil {
		panic(err)
	}

	if err = genFlag(appName); err != nil {
		panic(err)
	}

	if err = genConfig(appName); err != nil {
		panic(err)
	}

	if err = genCmd(appName); err != nil {
		panic(err)
	}
}
