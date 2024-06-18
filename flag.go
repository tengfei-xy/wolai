package main

import (
	"fmt"
	"os"

	log "github.com/tengfei-xy/go-log"
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
func (f *iflag) checkLogLevel() {
	log.SetLevelStr(f.loglevel)
	_, v := log.GetLevel()
	log.Infof("当前日志等级:%s", v)
}
