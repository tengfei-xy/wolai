package main

import (
	"fmt"
	"os"
)

type iflag struct {
	help       bool
	version    bool
	configFile string
	loglevel   string
}

func (f *iflag) checkHelp() {

	if f.help {
		description()
		os.Exit(0)
	}

}
func (f *iflag) checkVersion() {

	if f.version {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

}
