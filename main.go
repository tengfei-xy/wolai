package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/tengfei-xy/go-log"
)

const version string = "v1.2.2"

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
	flag.StringVar(&v.loglevel, "l", log.LEVELINFO, fmt.Sprintf("日志等级,可设置参数:%s(将输出到控制台)", strings.Join(log.GetLevelAll(), "、")))

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

	// 从配置文件中获取配置
	config, err = initConfig(f.configFile)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化日志系统：优先使用配置文件中的设置
	if config.Log.Level != "" || config.Log.OutputFile != "" {
		// 使用配置文件中的日志设置
		err = config.initLogger()
		if err != nil {
			log.Fatalf("初始化日志系统失败: %v", err)
		}
	} else {
		// 如果配置文件中没有设置，使用命令行参数
		log.SetLevelStr(f.loglevel)
	}
	_, v := log.GetLevel()
	log.Infof("当前日志等级:%s", v)

	// 在配置文件的保存地址上，增加时间
	config.addTimeBackupPath()

	// 获取用户ID用户所在的空间ID
	ui, err := getUserInfo()
	if err != nil {
		log.Fatal(err)
	}

	for _, ws := range ui.ws {

		// 根据用户所在的空间ID获取用户空间的基本信息
		ws.getBasicInfo()

		// 检查工作区是否应该被忽略
		if config.isIgnoreWorkspace(ws.name) {
			log.Warnf("根据配置文件忽略工作区:%s", ws.name)
			continue
		}

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

	// 关闭日志系统
	log.Close()
}
