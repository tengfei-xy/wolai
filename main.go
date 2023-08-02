package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

const version string = "v0.1"

func initMain(c config) error {
	if err := os.Mkdir(c.Save.newTargetPath, 0755); err != nil {
		return err
	}
	log.Infof("保存文件夹:%s", c.Save.newTargetPath)
	return nil
}
func parseFlag() bool {

	var exit bool = false
	helpText := flag.Bool("h", false, "查看帮助")
	versionText := flag.Bool("v", false, "查看版本")

	flag.Parse()
	if *helpText {
		fmt.Println("项目链接:https://github.com/tengfei-xy/wolai")
		exit = true

	}
	if *versionText {
		fmt.Printf("%s", version)
		exit = true
	}
	return exit
}
func main() {
	if exit := parseFlag(); exit {
		return
	}

	// 获取配置
	config, err := getConfig()
	if err != nil {
		log.Error(err)
		return
	}
	config.Save.newTargetPath = filepath.Join(config.Save.TargetPATH, timeGetChineseString())

	// 获取所有的总页面ID和名称
	pages, ok := getPagesList(config.Cookie)
	if !ok {
		return
	}

	// 初始化备份文件夹
	if err := initMain(config); err != nil {
		log.Error(err)
		return
	}

	for _, p := range pages {

		if tools.ListHasString(config.IgnorePageName, p.workSpacePageName) {
			log.Infof("忽略工作区页面 名称:%s", p.workSpacePageName)
			continue
		}

		reqJson, err := pageInfoToExportReqJsonAll(p)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Infof("正在导出%s", p.workSpacePageName)

		// 获取 导出完成的下载的url
		downloadUrl, err := exportMD(reqJson, config.Cookie)
		if err != nil {
			continue
		}

		// 设置下载链接和文件名
		p.url = downloadUrl
		u, _ := url.ParseRequestURI(p.url)
		p.filename = filepath.Base(u.Path)

		// 下载文件
		if err := tools.FileDownload(p.url, filepath.Join(config.newTargetPath, p.filename)); err != nil {
			log.Error(err)
			continue
		}
		log.Infof("下载成功 文件名:%s 链接:%s", p.filename, p.url)
	}

}
