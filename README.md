# Wolai 数据导出工具

![Go Version](https://img.shields.io/github/go-mod/go-version/tengfei-xy/wolai)
![License](https://img.shields.io/github/license/tengfei-xy/wolai)

一个用于导出 wolai（我来） 笔记的命令行工具，支持将所有文章导出为 Markdown 和 HTML 格式，防止数据丢失。

## ✨ 特性

- 🚀 支持批量导出所有 wolai 文章
- 📄 支持 Markdown 和 HTML 两种导出格式
- 🛡️ 数据备份保护，防止意外丢失
- ⚙️ 灵活的配置选项和忽略规则
- 📊 详细的导出进度和日志记录
- 🔧 支持多工作区和子空间管理

## 📋 使用场景

- **数据备份**：定期备份重要笔记内容
- **平台迁移**：从 wolai 迁移到其他笔记平台
- **离线查看**：本地保存文档，方便离线查阅
- **内容整理**：批量整理和归档笔记内容

## ⚠️ 重要说明

- **编码要求**：Windows 系统下配置文件需保存为 ANSI/GB2312 编码，其他系统使用 UTF-8
- **权限限制**：协作访客的工作区在导出时将得到空数据，因为这类工作区没有导出权限
- **性能参考**：约 2K 文档的页面导出耗时 - MD 格式：9 分钟+，HTML 格式：3 分钟+
- **技术实现**：由于官方暂未提供正式的导出 [API](https://www.wolai.com/wolai/7FB9PLeqZ1ni9FfD11WuUi)，本工具直接使用浏览器与后端的交互接口

## 🚀 快速开始

### 1. 下载可执行文件

从 [Releases](https://github.com/tengfei-xy/wolai/releases) 页面下载适合您操作系统的可执行文件。

### 2. 获取 Cookie

1. 在浏览器中登录 wolai
2. 打开浏览器开发者工具（F12）
3. 在 Network 标签页找到请求头中的 Cookie
4. 复制完整的 Cookie 字符串

### 3. 配置和运行

```bash
# 首次运行生成配置文件
./wolai -c config.yaml

# 编辑配置文件（参见下方配置说明）
vim config.yaml

# 运行导出
./wolai -c config.yaml
```

## ⚙️ 配置说明

配置文件 `config.yaml` 的详细说明：

### 基本配置

```yaml
# 🔑 必填：从浏览器登录后获取的 Cookie
cookie: "isg=xxx; token=xxx, wolai_client_id=xxx"

# 📁 必填：备份保存目录（建议使用绝对路径）
backupBackupDir: /Users/username/Documents/wolai

# 📄 导出格式（可选多个）
exportType:
  - html # HTML 格式，保持原版样式
  - md # Markdown 格式，便于编辑
```

### 日志配置

```yaml
log:
  # 日志级别：fatal, error, warn, info, debug, debug1, debug2, debug3
  level: info

  # 日志文件路径（为空则不输出到文件）
  outputFile: "debug.log"

  # 是否输出到控制台
  toConsole: true
```

### 忽略规则配置

支持灵活的忽略规则，可以跳过不需要导出的内容：

```yaml
# 忽略特定工作区、子空间或页面
# 💡 对于个人版本：子空间配置可选
# 💡 对于团队版本：需要明确指定子空间配置
ignore:
  - workspace: 工作区A
    subspace:
      - name: 子空间A1
        page:
          - name: "*" # 忽略所有页面
      - name: 子空间A2
        page:
          - name: 页面名1
          - name: 页面名2
  - workspace: 工作区B
    subspace:
      - name: 子空间B1
        page:
          - name: 页面名1
          - name: 页面名2
```

## 🛠️ 开发和编译

### 环境要求

- Go 1.21.4 或更高版本

### 编译步骤

```bash
# 克隆项目
git clone https://github.com/tengfei-xy/wolai.git

# 进入目录
cd wolai

# 编译
go build

# 或者直接运行
go run . -c config.yaml
```

### 跨平台编译

```bash
# Windows 64位
GOOS=windows GOARCH=amd64 go build -o wolai.exe

# Linux 64位
GOOS=linux GOARCH=amd64 go build -o wolai

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o wolai
```

## 📋 运行示例

以下是一个典型的运行日志示例：

```text
$ go run . -c config.yaml

2024-03-21 16:16:29 main.go:56 info  读取配置文件:config.yaml
2024-03-21 16:16:29 user.go:25 info  发现用户:6Wrpf1DKWh3LPzKwRzEDck
2024-03-21 16:16:30 workspace.go:96 warn  跳过工作区:xxx 原因:该用户是协作访客
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/
2024-03-21 16:16:30 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:13 html.go:74 info  下载成功 xxx
2024-03-21 16:23:54 main.go:129 info  导出结束!欢迎再次使用
```

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 👤 作者

- **作者**: [tengfei-xy](https://github.com/tengfei-xy)
- **项目主页**: [https://github.com/tengfei-xy/wolai](https://github.com/tengfei-xy/wolai)

## ⭐ 支持

如果这个项目对您有帮助，请给个 ⭐ Star！

---

> 💡 **提示**: 建议定期备份重要数据，避免因为平台变化造成数据丢失。
