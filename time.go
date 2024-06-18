package main

import (
	"math/rand"
	"time"
)

func timeGetChineseString() string {
	t := time.Now()
	return t.Format("2006年01月02日15点04分")
}
func createRand(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const pool string = "qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP12345678900"
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = pool[r.Intn(62)]
	}
	return string(bytes)
}
