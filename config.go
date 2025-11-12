package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
	"gopkg.in/yaml.v3"
)

func initConfig(config string) (Config, error) {
	var c Config
	f := getAppPath()

	file := filepath.Join(f, config)

	log.Infof("读取配置文件:%s", file)

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

	// 构建工作空间名称到索引的缓存
	c.buildWorkspaceCache()

	return c, c.check()
}

func (c *Config) initLogger() error {
	// 根据字符串日志级别获取对应的整数级别
	var levelInt int
	switch strings.ToLower(c.Log.Level) {
	case "fatal":
		levelInt = log.LEVELFATALINT
	case "error":
		levelInt = log.LEVELERRORINT
	case "warn":
		levelInt = log.LEVELWARNINT
	case "info":
		levelInt = log.LEVELINFOINT
	case "debug":
		levelInt = log.LEVELDEBUGINT
	case "debug1":
		levelInt = log.LEVELDEBUG1INT
	case "debug2":
		levelInt = log.LEVELDEBUG2INT
	case "debug3":
		levelInt = log.LEVELDEBUG3INT
	default:
		levelInt = log.LEVELINFOINT
	}

	// 创建日志配置
	logConfig := log.LogConfig{
		Level:      levelInt,
		OutputFile: c.Log.OutputFile,
		ToConsole:  c.Log.ToConsole,
	}

	// 初始化日志系统
	return log.Init(logConfig)
}

func (c *Config) check() error {
	if c.Cookie == "" {
		return fmt.Errorf("配置文件中的cookie值为空")
	}

	// 验证日志级别
	validLogLevels := []string{"fatal", "error", "warn", "info", "debug", "debug1", "debug2", "debug3"}
	if c.Log.Level == "" {
		c.Log.Level = "info" // 设置默认值
	} else {
		validLevel := false
		for _, level := range validLogLevels {
			if c.Log.Level == level {
				validLevel = true
				break
			}
		}
		if !validLevel {
			return fmt.Errorf("无效的日志级别: %s, 支持的级别为: %s", c.Log.Level, strings.Join(validLogLevels, ", "))
		}
	}

	// 设置默认的日志配置值
	// LogOutputFile 可以为空，表示只输出到控制台
	// LogToConsole 默认为 true

	_, err := os.Stat(c.BackupPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("路径不存在 %s", c.BackupPath)
	} else if err != nil {
		return fmt.Errorf("路径:%s 发生错误: %v ", c.BackupPath, err)
	}

	return nil
}

// buildWorkspaceCache 构建工作空间名称到索引的缓存映射
func (c *Config) buildWorkspaceCache() {
	c.workspaceIndexCache = make(map[string]int, len(c.Ignore))
	for i, workspace := range c.Ignore {
		c.workspaceIndexCache[workspace.Name] = i
	}
}

func (c *Config) getIgnoreWorkspace(ws string) int {
	// 如果缓存未初始化，先构建缓存
	if c.workspaceIndexCache == nil {
		c.buildWorkspaceCache()
	}

	if index, exists := c.workspaceIndexCache[ws]; exists {
		return index
	}
	return -1
}

// isIgnoreWorkspace 检查工作区是否应该被完全忽略
// 当工作区下所有子空间都配置为通配符 "*" 时，表示整个工作区都应该被忽略
func (c *Config) isIgnoreWorkspace(ws string) bool {
	wsIndex := c.getIgnoreWorkspace(ws)
	if wsIndex == -1 {
		return false
	}

	// 检查是否有子空间配置为通配符 "*"
	for _, subspace := range c.Ignore[wsIndex].Subspaces {
		if subspace.Name == "*" {
			return true
		}
	}

	return false
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

	// 快速检查边界条件
	if ws == -1 || ws >= len(c.Ignore) || len(subspaceName) == 0 {
		return false
	}

	// 在一次循环中检查通配符和精确匹配
	for _, subspace := range c.Ignore[ws].Subspaces {
		if subspace.Name == "*" {
			return true
		}
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
func (c *Config) addTimeBackupPath() {
	config.BackupPath = filepath.Join(config.BackupPath, timeGetChineseString())
}
func configGenerate(file string) error {
	var c Config
	c.BackupPath = getAppPath()
	c.Log.Level = "info"
	c.Log.OutputFile = ""
	c.Log.ToConsole = true
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
		return err
	}
	if err := tools.FileWrite(file, data); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println("对于团队版（家庭版），子空间是必须的，需要填写此参数，subspace项可以有多个")
	fmt.Println("对于个人版，子空间是默认第一个的，无需（或任意）填写此参数")
	return fmt.Errorf("已创建新配置文件,请修改后重新运行程序 位置:%s", file)
}

type Config struct {
	Cookie string `yaml:"cookie"`
	Log    Log    `yaml:"log"`

	// 备份路径
	BackupPath string      `yaml:"backupBackupDir"`
	ExportType []string    `yaml:"exportType"`
	Ignore     []Workspace `yaml:"ignore"`

	// 缓存：工作空间名称到索引的映射
	workspaceIndexCache map[string]int
}
type Workspace struct {
	Name      string     `yaml:"workspace"`
	Subspaces []Subspace `yaml:"subspace"`
}
type Log struct {
	Level      string `yaml:"level"`
	OutputFile string `yaml:"outputFile"`
	ToConsole  bool   `yaml:"toConsole"`
}
type Subspace struct {
	Name  string `yaml:"name"`
	Pages []Page `yaml:"page"`
}

type Page struct {
	Name string `yaml:"name"`
}
