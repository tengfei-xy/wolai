package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/tengfei-xy/go-log"
)

const version string = "v1.2.1"

var config Config

func description() {
	fmt.Printf(`                  _         _ 
__      __  ___  | |  __ _ (_)
\ \ /\ / / / _ \ | | / _`+"`"+` || |
 \ V  V / | (_) || || (_| || |   程序版本: %s
  \_/\_/   \___/ |_| \__,_||_|   项目链接: https://github.com/tengfei-xy/wolai
                              %s`, version, "\n")

}
func mkdir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil && err != os.ErrExist {
		return err
	}
	log.Infof("创建文件夹 路径:%s", path)
	return nil
}

func initFlag() iflag {
	var v iflag
	flag.BoolVar(&v.help, "h", false, "查看帮助")
	flag.BoolVar(&v.version, "v", false, "查看版本")
	flag.StringVar(&v.configFile, "c", "config.yaml", "指定配置文件")
	flag.StringVar(&v.loglevel, "l", log.LEVELINFO, fmt.Sprintf("日志等级,可设置参数:%s", strings.Join(log.GetLevelAll(), "、")))

	flag.Parse()
	return v
}
func initTilte() {
	description()
}
func main() {
	var err error
	initTilte()

	f := initFlag()
	f.checkHelp()
	f.checkVersion()
	f.checkLogLevel()

	// 从配置文件中获取配置
	config, err = initConfig(f.configFile)
	if err != nil {
		log.Fatal(err)
	}

	// 在配置文件的保存地址上，增加时间
	config.addTimeBackupPath()

	// 获取用户ID用户所在的空间ID
	ui, err := getUserInfo()
	if err != nil {
		log.Fatal(err)
	}

	// // 获取官方API的工作区结构
	// ws, err := getWorkSpaceStruct()
	// if err != nil {
	// 	panic(err)
	// }

	// // 将官方API的工作区结构转化为重要字段的结构体 workspaceInfo
	// wsInfos := ws.getWorkspaceInfo(ui.userid)

	for _, ws := range ui.ws {

		// 根据用户所在的空间ID获取用户空间的基本信息
		ws.getBasicInfo()

		// 一个子空间的处理方式
		if ws.isDefaultSubWorkspace() {

			// 获取默认子空间的信息
			ws.getDefaultSubspace()
		} else {
			// 设定结构体长度并子空间的获取ID
			ws.getTeamSubspace()

			// 获取名称
			ws.getTermPagesMain()
		}

		// 进行导出MD
		ws.exportMDMain()

		// 进行导出HTML
		ws.exportHTMLMain()
	}

	log.Info("导出结束!欢迎再次使用")

}
