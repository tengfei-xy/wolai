# wolai
wolai的一些便捷功能，支持导出所有文章（怕官方跑路或丢失数据）





## 说明

- windows的配置文件需要保存为ANSI编码，其他系统使用utf-8
- 微信：SXL--LP（请备注Github哦）
- 后来遇到了这个应用数据丢失的问题，且无法恢复，历史记录不存在。然而md格式不太具有美观，于是新版本里添加html方便为迁移到平台提供原版的画面。
- wolai上，我有2k个文档，耗时了9分钟完成
- 由于官方目前没有给出导出md的正式[API](https://www.wolai.com/wolai/7FB9PLeqZ1ni9FfD11WuUi)，所以是直接用了浏览器与后端直接交互接口
- 会跳过作为协作访客的工作区，因为这类工作区没有导出的权限





## 使用方法

[下载地址](https://github.com/tengfei-xy/wolai/releases)

1. 首次运行将生成config.yaml

2. 修改配置文件后，重新运行主程序

3. 使用方法

```
./wolai -c config.yaml
```



## 配置文件说明

```yaml
# 请从浏览器登录自己的个人空间后获取
cookie: "isg=xxx; token=xxx, wolai_client_id=xxx"

# 作为保存的目标文件夹
# 建议为绝对路径
backupBackupDir: /Users/melta/Documents/wolai

# 设定导出类型
# 大小写无关，支持值有html、md或markdown
exportType: ["html","md"]

# 对于团队版（家庭版），子空间是必须的，需要填写此参数，subspace项可以有多个
# 对于个人版，子空间是默认第一个的，无需（或任意）填写此参数
ignore:
    - workspace: 工作区名
      subspace:
        - name: 子空间名
          page:
            - name: "*" # 表示忽略所有
        - name: 子空间名
          page:
            - name: 页面名
            - name: 页面名
    - workspace: 工作区名
      subspace:
        - name: 子空间名
          page:
            - name: 公司
            - name: 页面名
        - name: 子空间名
          page:
            - name: 页面名
            - name: 页面名

```



## 项目编译

编译（版本：1.19.6）

```shell
go clone https://github.com/tengfei-xy/wolai.git
cd wolai
go build
```





## 日志

```
XiaoXiaoFeiTongXue  wolai git:(main)   go run . -c config.yaml
2024-03-21 16:16:29 main.go:56 info  读取配置文件:config.yaml
2024-03-21 16:16:29 user.go:25 info  发现用户:6Wrpf1DKWh3LPzKwRzEDck
2024-03-21 16:16:30 workspace.go:96 warn  跳过工作区:xxx 原因:该用户是协作访客
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:公开分享 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:梦娜·学不进·想自莎 子空间:公开分享 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 workspace.go:292 info  发现工作区:xxx 页面:xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/xxx
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/梦娜·学不进·想自莎/公开分享
2024-03-21 16:16:30 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2024年03月21日16点16分/html/xxx
2024-03-21 16:16:30 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:33 html.go:74 info  下载成功 xxx
2024-03-21 16:16:33 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:48 html.go:74 info  下载成功 xxx
2024-03-21 16:16:48 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:51 html.go:74 info  下载成功 xxx
2024-03-21 16:16:51 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:53 html.go:74 info  下载成功 xxx
2024-03-21 16:16:53 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:55 html.go:74 info  下载成功 xxx
2024-03-21 16:16:55 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:16:58 html.go:74 info  下载成功 xxx
2024-03-21 16:16:58 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:00 html.go:74 info  下载成功 xxx
2024-03-21 16:17:00 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:03 html.go:74 info  下载成功 xxx
2024-03-21 16:17:03 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:06 html.go:74 info  下载成功 xxx
2024-03-21 16:17:06 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:10 html.go:74 info  下载成功 xxx
2024-03-21 16:17:10 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:14 html.go:74 info  下载成功 xxx
2024-03-21 16:17:14 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:17 html.go:74 info  下载成功 xxx
2024-03-21 16:17:17 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:22 html.go:74 info  下载成功 xxx
2024-03-21 16:17:22 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:25 html.go:74 info  下载成功 xxx
2024-03-21 16:17:25 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:28 html.go:74 info  下载成功 xxx
2024-03-21 16:17:28 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:31 html.go:74 info  下载成功 xxx
2024-03-21 16:17:31 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:33 html.go:74 info  下载成功 xxx
2024-03-21 16:17:33 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:42 html.go:74 info  下载成功 xxx
2024-03-21 16:17:42 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:45 html.go:74 info  下载成功 xxx
2024-03-21 16:17:45 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:17:50 html.go:74 info  下载成功 xxx
2024-03-21 16:17:50 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:00 html.go:74 info  下载成功 xxx
2024-03-21 16:18:00 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:03 html.go:74 info  下载成功 xxx
2024-03-21 16:18:03 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:06 html.go:74 info  下载成功 xxx
2024-03-21 16:18:06 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:10 html.go:74 info  下载成功 xxx
2024-03-21 16:18:10 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:13 html.go:74 info  下载成功 xxx
2024-03-21 16:18:13 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:18:16 html.go:74 info  下载成功 xxx
2024-03-21 16:18:16 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:21:21 html.go:74 info  下载成功 xxx
2024-03-21 16:21:21 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:21:25 html.go:74 info  下载成功 xxx
2024-03-21 16:21:25 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:21:28 html.go:74 info  下载成功 xxx
2024-03-21 16:21:28 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:22:17 html.go:74 info  下载成功 xxx
2024-03-21 16:22:17 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:22:22 html.go:74 info  下载成功 xxx
2024-03-21 16:22:22 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:22:26 html.go:74 info  下载成功 xxx
2024-03-21 16:22:26 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:22:31 html.go:74 info  下载成功 xxx
2024-03-21 16:22:31 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:13 html.go:74 info  下载成功 xxx
2024-03-21 16:23:13 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:17 html.go:74 info  下载成功 xxx
2024-03-21 16:23:17 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:19 html.go:74 info  下载成功 xxx
2024-03-21 16:23:19 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:23 html.go:74 info  下载成功 xxx
2024-03-21 16:23:23 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:27 html.go:74 info  下载成功 xxx
2024-03-21 16:23:27 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:xxx 页面:xxx
2024-03-21 16:23:29 html.go:74 info  下载成功 xxx
2024-03-21 16:23:29 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:公开分享 页面:xxx
2024-03-21 16:23:32 html.go:74 info  下载成功 xxx
2024-03-21 16:23:32 export.go:18 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:公开分享 页面:xxx
2024-03-21 16:23:54 html.go:74 info  下载成功 xxx
2024-03-21 16:23:54 html.go:23 warn  忽略导出 工作区:xxx 子空间:
2024-03-21 16:23:54 main.go:129 info  导出结束!欢迎再次使用
```


