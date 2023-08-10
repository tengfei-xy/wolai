package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

func exportMain(dmd downloadMD) {
	var ignorePages []string
	spaceName := dmd.workSpaceName

	// 判断是否忽略空间
	for _, igPageName := range dmd.workSpacePageName {
		if igPageName == "*" {
			log.Infof("忽略 空间名:%s", dmd.workSpaceName)
			return
		}
	}

	// 获取当前需要忽略的页面列表
	for _, space := range config.Ignore {
		if dmd.workSpaceName == space.SpaceName {
			ignorePages = space.Page
			break
		}
	}

	// 设定导出的空间文件夹
	spaceFolder := filepath.Join(config.Save.newTargetPath, spaceName)
	if err := os.Mkdir(spaceFolder, 0755); err != nil {
		log.Error(err)
		return
	}

	for i, id := range dmd.workSpacePageID {
		name := dmd.workSpacePageName[i]

		// 忽略页面
		if tools.ListHasString(ignorePages, name) {
			log.Infof("忽略 空间名:%s 页面名称:%s", spaceName, name)
			continue
		}

		reqJson, err := exportReqJsonAll(id, name)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Infof("导出 空间名:%s 页面名:%s", spaceName, name)

		// 获取 导出完成的下载的url
		fileURL, err := exportMD(reqJson, config.Cookie)
		if err != nil {
			continue
		}

		// 设置下载链接和文件名
		u, _ := url.ParseRequestURI(fileURL)
		filename := filepath.Base(u.Path)

		// 下载文件
		if err := tools.FileDownload(fileURL, filepath.Join(spaceFolder, filename)); err != nil {
			log.Error(err)
			continue
		}
		log.Infof("下载成功 文件名:%s 链接:%s", filename, fileURL)
	}

}
func exportReqJsonAll(id, name string) ([]byte, error) {
	var e exportUP
	e.PageID = id
	e.PageTitle = name
	e.Options.RecoverTree = true
	e.Options.GenerateToc = "all"
	e.Options.IncludeSubPage = true
	return json.Marshal(e)
}
func exportMD(data []byte, cookie string) (string, error) {
	d, ok := exportMarkdownData(data, cookie)
	if !ok {
		return "", fmt.Errorf("获取失败")
	}
	return exportMarkdownDeal(d)
}
func exportMarkdownData(data []byte, cookie string) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/exportMarkdown`, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("host", `api.wolai.com`)
	req.Header.Set("Accept-Encoding", `gzip, deflate, br`)
	req.Header.Set("Origin", `https://www.wolai.com`)
	req.Header.Set("Sec-Fetch-Dest", `empty`)
	req.Header.Set("Sec-Fetch-Site", `same-site`)
	req.Header.Set("Accept-Language", `en-US,en;q=0.9`)
	req.Header.Set("Accept", `application/json, text/plain, */*`)
	req.Header.Set("Content-Type", ` application/json`)
	req.Header.Set("Sec-Fetch-Mode", `cors`)
	req.Header.Set("wolai-os-platform", `mac`)
	req.Header.Set("x-client-timezone", `Asia/Shanghai`)
	req.Header.Set("wolai-app-version", `1.1.2-3`)
	req.Header.Set("wolai-client-platform", `web`)
	req.Header.Set("x-client-timeoffset", `-480`)
	req.Header.Set("wolai-client-version", ``)
	req.Header.Set("reqId", "7NYkp7BR1wdJm1oXFRevfJ")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("状态码:%d", resp.StatusCode)
		return nil, false
	}

	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return nil, false
	}
	return resp_data, true
}
func exportMarkdownDeal(data []byte) (string, error) {
	var e exportStruct
	if err := json.Unmarshal(data, &e); err != nil {
		log.Error(err)
		return "", err
	}
	if e.Code != 1000 {
		err := fmt.Errorf("%s", e.Message)
		log.Error(err)
		return "", err
	}
	return e.Data, nil
}

type exportUP struct {
	PageID    string          `json:"pageId"`
	PageTitle string          `json:"pageTitle"`
	Options   exportUPOptions `json:"options"`
}
type exportUPOptions struct {
	RecoverTree    bool   `json:"recoverTree"`
	GenerateToc    string `json:"generateToc"`
	IncludeSubPage bool   `json:"includeSubPage"`
}

type exportStruct struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	Title     string `json:"title"`
	RequestID string `json:"requestId"`
}
