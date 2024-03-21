package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	tools "github.com/tengfei-xy/go-tools"
	"gopkg.in/yaml.v3"
)

func initConfig(config string) (Config, error) {
	var c Config
	f := getAppPath()

	file := filepath.Join(f, config)

	if !tools.FileExist(file) {
		return Config{}, configGenerate(file)
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

	return c, c.check()
}

func (c *Config) check() error {
	if c.Cookie == "" {
		return fmt.Errorf("配置文件中的cookie值为空")
	}
	_, err := os.Stat(c.BackupPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("路径不存在 %s", c.BackupPath)
	} else if err != nil {
		return fmt.Errorf("路径:%s 发生错误: %v ", c.BackupPath, err)
	}

	return nil
}
func (c *Config) getIgnoreWorkspace(ws string) int {
	for i, j := range c.Ignore {
		if j.Name == ws {
			return i
		}
	}
	return -1
}
func (c *Config) getIgnoreSubspace(freePlan bool, ws int, sp string) int {
	if ws == -1 {
		return -1
	}
	if freePlan {
		return 0
	}
	for i, j := range c.Ignore[ws].Subspaces {
		if j.Name == sp {
			return i
		}
	}
	return -1
}
func (c *Config) isIgnoreSubspace(ws int, subspaceName string) bool {
	if ws == -1 {
		return false
	}
	for _, subspace := range c.Ignore[ws].Subspaces {
		if subspace.Name == "*" {
			return true
		}
	}

	for _, subspace := range c.Ignore[ws].Subspaces {
		if subspace.Name == subspaceName {
			return true
		}
	}
	return false
}
func (c *Config) isIgnorePage(ws int, sb int, pageName string) bool {
	if sb == -1 || ws == -1 {
		return false
	}
	for _, page := range c.Ignore[ws].Subspaces[sb].Pages {
		if page.Name == "*" {
			return true
		}
	}

	for _, page := range c.Ignore[ws].Subspaces[sb].Pages {
		if page.Name == pageName {
			return true
		}
	}
	return false
}
func (c *Config) hasHtml() bool {
	for _, v := range c.ExportType {
		if strings.ToLower(v) == "html" {
			return true
		}
	}
	return false

}
func (c *Config) hasMarkdown() bool {

	for _, v := range c.ExportType {
		v = strings.ToLower(v)
		if v == "markdown" || v == "md" {
			return true
		}
	}
	return false

}
func configGenerate(file string) error {
	var c Config
	c.BackupPath = getAppPath()
	c.Ignore = make([]Workspace, 2)
	c.ExportType = []string{"html", "markdown"}
	for i := range c.Ignore {
		c.Ignore[i].Name = "工作区名"
		c.Ignore[i].Subspaces = make([]Subspace, 2)
		for k := range c.Ignore[i].Subspaces {
			c.Ignore[i].Subspaces[k].Name = "子空间名"
			c.Ignore[i].Subspaces[k].Pages = make([]Page, 2)
			for q := range c.Ignore[i].Subspaces[k].Pages {
				c.Ignore[i].Subspaces[k].Pages[q].Name = "页面名"
			}
		}
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}
	if err := tools.FileWrite(file, data); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println("对于团队版（家庭版），子空间是必须的，需要填写此参数，subspace项可以有多个")
	fmt.Println("对于个人版，子空间是默认第一个的，无需（或任意）填写此参数")
	return fmt.Errorf("已创建新配置文件,请修改后重新运行程序 位置:%s", file)
}

type Config struct {
	Cookie     string      `yaml:"cookie"`
	BackupPath string      `yaml:"backupBackupDir"`
	ExportType []string    `yaml:"exportType"`
	Ignore     []Workspace `yaml:"ignore"`
}
type Workspace struct {
	Name      string     `yaml:"workspace"`
	Subspaces []Subspace `yaml:"subspace"`
}

type Subspace struct {
	Name  string `yaml:"name"`
	Pages []Page `yaml:"page"`
}

type Page struct {
	Name string `yaml:"name"`
}
