package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(p downloadMD) error {

	url := p.url
	target := p.filename

	res, err := send_get_request(url)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(getCurrentPath()+target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err2 := f.Write(res); !(err2 == nil && err2 != io.EOF) {
		return err
	}
	return nil
}

func send_get_request(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("url:%s StatusCode:%d", link, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
