package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

const version string = "v0.2"

var config Config

func project() {
	fmt.Printf("程序版本:%s\n", version)
	fmt.Println("项目链接:https://github.com/tengfei-xy/wolai")

}
func initMain(c Config) error {
	if err := os.Mkdir(c.Save.newTargetPath, 0755); err != nil {
		return err
	}
	log.Infof("创建 保存目标文件夹:%s", c.Save.newTargetPath)
	return nil
}
func parseFlag() bool {

	var exit bool = false
	helpText := flag.Bool("h", false, "查看帮助")
	versionText := flag.Bool("v", false, "查看版本")

	flag.Parse()
	if *helpText {
		project()
		exit = true

	}
	if *versionText {
		fmt.Printf("%s", version)
		exit = true
	}
	return exit
}
func main() {
	var err error

	if exit := parseFlag(); exit {
		return
	}

	// 获取配置
	config, err = getConfig()
	if err != nil {
		log.Error(err)
		tools.Delay(5)
		return
	}

	config.Save.newTargetPath = filepath.Join(config.Save.TargetPATH, timeGetChineseString())

	// 获取所有的总页面ID和名称
	workspace, ok := getPagesList(config.Cookie)
	if !ok {
		tools.Delay(5)
		return
	}

	// 初始化备份文件夹
	if err := initMain(config); err != nil {
		log.Error(err)
		tools.Delay(5)
		return
	}

	// 开始导出
	for _, space := range workspace {
		exportMain(space)
	}

	tools.Delay(5)
}
