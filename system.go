package main

import (
	"os"
	"path/filepath"
	"runtime"
)

// 作用:返回程序的文件的所在路径
func getAppPath() string {
	var abWorkPath string
	if runtime.GOOS == "windows" {
		exePath, _ := os.Executable()
		abWorkPath, _ = filepath.EvalSymlinks(filepath.Dir(exePath))
	} else {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abWorkPath = filepath.Dir(filename)
		}
	}
	return abWorkPath
}
