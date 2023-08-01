package main

import (
	"net/url"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
)

func main() {

	cookie := ``
	reqId := `7NYkp7BR1wdJm1oXFRevfJ`

	// 获取所有的总页面ID
	pages, ok := getPagesList(cookie, reqId)
	if !ok {
		return
	}

	for _, p := range pages {
		reqJson, err := pageInfoToExportReqJsonAll(p)
		if err != nil {
			log.Error(err)
			continue
		}

		// 这一步,仅仅获取 url 而已
		downloadUrl, err := exportMD(reqJson, cookie, reqId)
		if err != nil {
			continue
		}
		p.url = downloadUrl
		u, _ := url.ParseRequestURI(p.url)
		p.filename = filepath.Base(u.Path)

		if err := downloadFile(p); err != nil {
			continue
		}
		log.Infof("下载成功 文件名:%s 链接:%s", p.filename, p.url)
	}

}
