package main

import "time"

func timeGetChineseString() string {
	t := time.Now()
	return t.Format("2006年01月02日15点04分")
}
