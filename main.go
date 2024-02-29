package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/tengfei-xy/go-log"
)

const version string = "v0.3.0"

var config Config

func project() {
	fmt.Printf("程序版本:%s\n", version)
	fmt.Println("项目链接:https://github.com/tengfei-xy/wolai")

}
func mkdir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil && err != os.ErrExist {
		return err
	}
	log.Infof("创建 保存目标文件夹:%s", path)
	return nil
}

func initFalg() (bool, string) {

	var exit bool = false
	helpText := flag.Bool("h", false, "查看帮助")
	versionText := flag.Bool("v", false, "查看版本")
	configFileText := flag.String("c", "config.yaml", "指定配置文件")

	flag.Parse()
	if *helpText {
		project()
		exit = true

	}
	if *versionText {
		fmt.Printf("%s", version)
		exit = true
	}
	return exit, *configFileText
}
func main() {
	var err error

	exit, f := initFalg()
	if exit {
		return
	}
	log.Infof("读取配置文件:%s", f)

	// 获取配置
	config, err = initConfig(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	userid, err := getUserID()
	if err != nil {
		panic(err)
	}

	// 获取官方API的工作区结构
	ws, err := getWorkSpaceStruct()
	if err != nil {
		panic(err)
	}

	// 将官方API的工作区结构转化为重要字段的结构体 workspaceInfo
	wsInfos := ws.getWorkspaceInfo(userid)

	// 获取子空间
	for i := range wsInfos {

		if wsInfos[i].id == "" {
			continue
		}

		// 免费版
		if wsInfos[i].is_free_plan() {
			if err := wsInfos[i].getDefaultSubspace(); err != nil {
				panic(err)
			}
			continue
		}

		// 多人工作区 设定结构体长度并子空间的获取ID
		if err := wsInfos[i].getTeamSubspace(); err != nil {
			panic(err)
		}
		// 多人工作区 获取名称
		if err := wsInfos[i].getTermPagesMain(); err != nil {
			panic(err)
		}

	}

	// 输出将被导出的页面
	for _, wsInfo := range wsInfos {
		wsInfo.output()
	}

	config.BackupPath = filepath.Join(config.BackupPath, timeGetChineseString())
	for _, wsInfo := range wsInfos {
		if err := wsInfo.mkdirBackupFolder(); err != nil {
			panic(err)
		}
	}

	// 开始导出
	for _, wsInfo := range wsInfos {
		wsInfo.exportMain()
	}

	log.Info("导出结束!欢迎再次使用")

}
