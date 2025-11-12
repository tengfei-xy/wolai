package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

func (wsInfo *workspaceInfo) exportHTMLMain() {
	if !config.hasHtml() {
		log.Debugf("忽略导出HTML 名称:%s", wsInfo.name)
		return
	}

	// 使用公共的导出处理逻辑
	wsInfo.processExportWithIgnore(wsInfo.exportHtmlSingle)
}

func (wsInfo *workspaceInfo) exportHtmlSingle(subspaceName, pageId, pageName string) error {
	var e exportHtmlUpJson
	e.PageID = pageId
	e.PageTitle = pageName
	e.Options.IncludeSubPage = true

	reqJson, err := json.Marshal(e)
	if err != nil {
		return err
	}

	// 获取 导出完成的下载的url
	fileURL, err := exportDealHtml(reqJson)
	if err != nil {
		return err
	}

	// 设置下载链接和文件名
	u, _ := url.ParseRequestURI(fileURL)
	filename := filepath.Base(u.Path)
	var savePath string = filepath.Join(config.BackupPath, "html", wsInfo.name, subspaceName)
	if wsInfo.isDefaultSubWorkspace() {
		savePath = filepath.Join(config.BackupPath, "html", wsInfo.name)
	}

	// 创建备份文件夹
	if err := mkdir(savePath); err != nil {
		log.Fatal(err)
	}

	filename = filepath.Join(savePath, filename)

	// 下载文件
	if err := tools.FileDownload(fileURL, filename); err != nil {
		return err
	}
	log.Infof("下载成功 格式:html 保存路径:%s 链接:%s", filename, fileURL)
	return nil
}
func exportDealHtml(data []byte) (string, error) {
	// 延迟时间序列：1, 2, 5, 10, 15, 20分钟
	delayMinutes := []int{1, 2, 5, 10, 15, 20}
	retryCount := 0

	for {
		d, ok := exportHtmlHtml(data)
		if !ok {
			return "", fmt.Errorf("获取失败")
		}

		s, err, isRetry := exportHtmlDeal(d)
		if !isRetry {
			// 成功或者不需要重试的错误，直接返回
			return s, err
		}

		// 需要重试
		if retryCount >= len(delayMinutes) {
			// 超过最大重试次数，使用最后一个延迟时间
			delayTime := delayMinutes[len(delayMinutes)-1]
			log.Infof("重试第 %d 次，等待 %d 分钟", retryCount+1, delayTime)
			time.Sleep(time.Duration(delayTime) * time.Minute)
		} else {
			delayTime := delayMinutes[retryCount]
			log.Infof("重试第 %d 次，等待 %d 分钟", retryCount+1, delayTime)
			time.Sleep(time.Duration(delayTime) * time.Minute)
		}

		retryCount++
	}
}
func exportHtmlHtml(data []byte) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/exportHtml`, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", config.Cookie)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("host", `api.wolai.com`)
	req.Header.Set("Origin", `https://www.wolai.com`)
	req.Header.Set("Sec-Fetch-Dest", `empty`)
	req.Header.Set("Sec-Fetch-Site", `same-site`)
	req.Header.Set("Accept-Language", `en-US,en;q=0.9`)
	req.Header.Set("Accept", `application/json, text/plain, */*`)
	req.Header.Set("Content-Type", ` application/json`)
	req.Header.Set("Sec-Fetch-Mode", `cors`)
	req.Header.Set("wolai-os-platform", `mac`)
	req.Header.Set("x-client-timezone", `Asia/Shanghai`)
	req.Header.Set("wolai-app-version", `1.2.2-4`)
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

func exportHtmlDeal(data []byte) (string, error, bool) {
	var e exportStruct
	if err := json.Unmarshal(data, &e); err != nil {
		log.Error(err)
		return "", err, false
	}
	if e.Code != 1000 {
		err := fmt.Errorf("%s", e.Message)
		if err.Error() == "导出频率过快，请稍后再试" {
			return "", err, true
		}
		log.Error(err)
		return "", err, false
	}
	return e.Data, nil, false
}

type exportHtmlUpJson struct {
	PageID    string  `json:"pageId"`
	PageTitle string  `json:"pageTitle"`
	Options   Options `json:"options"`
}
type Options struct {
	IncludeSubPage bool `json:"includeSubPage"`
}
