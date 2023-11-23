package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

func (wsInfo *workspaceInfo) outputIgnore(ws, sb, page string) {
	if wsInfo.is_free_plan() {
		log.Warnf("忽略导出 工作区:%s 页面:%s", ws, page)
	} else {
		log.Warnf("忽略导出 工作区:%s 子空间:%s 页面:%s", ws, sb, page)
	}
}
func (wsInfo *workspaceInfo) outputExport(ws, sb, page string) {
	if wsInfo.is_free_plan() {
		log.Infof("开始导出 工作区:%s 页面:%s", ws, page)
	} else {
		log.Infof("开始导出 工作区:%s 子空间:%s 页面:%s", ws, sb, page)
	}
}
func (wsInfo *workspaceInfo) exportMain() {

	// 获取忽略的工作区序号
	wsSeq := config.getIgnoreWorkspace(wsInfo.name)

	for _, subspace := range wsInfo.subspace {
		if config.isIgnoreSubspace(wsSeq, subspace.name) {
			log.Warnf("忽略导出 工作区:%s 子空间:%s", wsInfo.name, subspace.name)
			continue
		}
		spSeq := config.getIgnoreSubspace(wsInfo.is_free_plan(), wsSeq, subspace.name)

		for _, page := range subspace.pages {
			if config.isIgnorePage(wsSeq, spSeq, page.name) {
				wsInfo.outputIgnore(wsInfo.name, subspace.name, page.name)
				continue
			}
			wsInfo.outputExport(wsInfo.name, subspace.name, page.name)
			err := wsInfo.exportSingle(subspace.name, page.id, page.name)
			if err != nil {
				log.Error(err)
			}
		}

	}

}
func (wsInfo *workspaceInfo) exportSingle(subspaceName, pageId, pageName string) error {
	var e exportUpJson
	e.PageID = pageId
	e.PageTitle = pageName
	e.Options.RecoverTree = true
	e.Options.GenerateToc = "all"
	e.Options.IncludeSubPage = true

	reqJson, err := json.Marshal(e)
	if err != nil {
		return err
	}

	// 获取 导出完成的下载的url
	fileURL, err := exportMD(reqJson)
	if err != nil {
		return err
	}

	// 设置下载链接和文件名
	u, _ := url.ParseRequestURI(fileURL)
	filename := filepath.Base(u.Path)

	if wsInfo.is_free_plan() {
		filename = filepath.Join(config.BackupPath, wsInfo.name, filename)
	} else {
		filename = filepath.Join(config.BackupPath, wsInfo.name, subspaceName, filename)
	}

	// 下载文件
	if err := tools.FileDownload(fileURL, filename); err != nil {
		return err
	}
	log.Infof("下载成功 保存路径:%s 链接:%s", filename, fileURL)
	return nil
}

func exportMD(data []byte) (string, error) {
	d, ok := exportMarkdownHtml(data)
	if !ok {
		return "", fmt.Errorf("获取失败")
	}
	return exportMarkdownDeal(d)
}
func exportMarkdownHtml(data []byte) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/exportMarkdown`, bytes.NewReader(data))
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
	req.Header.Set("wolai-app-version", `1.2.0-11`)
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

type exportUpJson struct {
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
