package main

import (
	"fmt"
	"path/filepath"

	tools "github.com/tengfei-xy/go-tools"
	"gopkg.in/yaml.v3"
)

func getConfig() (config, error) {
	var c config
	f := getAppPath()

	file := filepath.Join(f, "config.yaml")

	if !tools.FileExist(file) {
		return config{}, configGenerate(file + ".tmp")
	}

	data, err := tools.FileRead(file)
	if err != nil {
		return config{}, err
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return config{}, err
	}

	return c, configOK(c)
}
func configOK(c config) error {
	if c.Cookie == "" {
		return fmt.Errorf("配置文件中的cookie值为空")
	}

	return nil
}
func configGenerate(file string) error {
	var c config
	c.TargetPATH = getAppPath()

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}
	if err := tools.FileWrite(file, data); err != nil {
		return fmt.Errorf("%s", err)
	}
	return fmt.Errorf("已创建新配置文件,请修改后重新运行程序 位置:%s", file)

}

type config struct {
	Login `yaml:"login"`
	Save  `yaml:"save"`
	Pages `yaml:"ignore"`
}
type Login struct {
	Cookie string `yaml:"cookie"`
}
type Save struct {
	TargetPATH    string `yaml:"targetpath"`
	newTargetPath string
}
type Pages struct {
	IgnorePageName []string `yaml:"ignorePageName"`
}
