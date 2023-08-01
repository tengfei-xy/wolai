package main

import (
	"os"
	"path/filepath"
	"runtime"
)

var sh string

func getCurrentPath() string {
	var abWorkPath string
	if runtime.GOOS == "windows" {
		sh = "\\"
		exePath, err := os.Executable()
		if err != nil {
			return `.\`
		}
		abWorkPath, _ = filepath.EvalSymlinks(filepath.Dir(exePath))
	} else {
		sh = "/"
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abWorkPath = filepath.Dir(filename)
		}
	}
	return abWorkPath

}
