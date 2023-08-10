package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	tools "github.com/tengfei-xy/go-tools"
	"gopkg.in/yaml.v3"
)

func getConfig() (Config, error) {
	var c Config
	f := getAppPath()

	file := filepath.Join(f, "config.yaml")

	if !tools.FileExist(file) {
		return Config{}, configGenerate(file + ".tmp")
	}

	data, err := tools.FileRead(file)
	if err != nil {
		return Config{}, err
	}

	if runtime.GOOS == "windows" {
		data, err = tools.StringGBKToUTF_8(data)
		if err != nil {
			return Config{}, err
		}
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return Config{}, err
	}

	return c, configOK(c)
}

func configOK(c Config) error {
	if c.Cookie == "" {
		return fmt.Errorf("配置文件中的cookie值为空")
	}

	return nil
}
func configGenerate(file string) error {
	var c Config
	c.TargetPATH = getAppPath()
	c.Ignore = make([]Space, 2)
	c.Ignore[0].SpaceName = "个人空间名"
	c.Ignore[1].SpaceName = "个人空间名"
	c.Ignore[0].Page = make([]string, 1)
	c.Ignore[1].Page = make([]string, 1)
	c.Ignore[0].Page[0] = "页面名"
	c.Ignore[1].Page[0] = "页面名"

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}
	if err := tools.FileWrite(file, data); err != nil {
		return fmt.Errorf("%s", err)
	}
	return fmt.Errorf("已创建新配置文件,请修改后重新运行程序 位置:%s", file)
}

type Config struct {
	Login  `yaml:"login"`
	Save   `yaml:"save"`
	Ignore []Space `yaml:"ignore"`
}
type Login struct {
	Cookie string `yaml:"cookie"`
}
type Save struct {
	TargetPATH    string `yaml:"targetpath"`
	newTargetPath string
}
type Space struct {
	SpaceName string   `yaml:"spaceName"`
	Page      []string `yaml:"pageName"`
}
