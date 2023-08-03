# wolai
wolai的一些便捷功能，支持导出所有文章（怕官方跑路哈哈哈哈）



## 使用方法

[下载地址](https://github.com/tengfei-xy/wolai/releases)

1.首次运行将生成config.yaml.tmp

2.修改配置文件后并重命名为config.yaml

3.重新运行主程序



## 配置文件说明

```
login:
    # 请从浏览器登录自己的个人空间后获取
    cookie: "isg=xxx; token=xxx, wolai_client_id=xxx"
save:
		# 作为保存的目标文件夹
    targetpath: "/Users/melta/Documents/wolai"
ignore:
		# 忽略的页面名称
    ignorePageName: ["计算机","xxx"]
```



## 系统说明

1. windows的配置文件需要保存为ANSI编码，其他系统使用utf-8



## 主要功能-导出全部md

> 注：由于官方目前没有给出导出md的正式[API](https://www.wolai.com/wolai/7FB9PLeqZ1ni9FfD11WuUi)，所以是直接用了浏览器与后端直接交互接口

使用方法

```
./wolai
```

日志

```
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:密码学
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:其他
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:哲学
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:政治
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:英语
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:计算机
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:物理学
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:人生
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:历史
2023-08-02 14:12:14 pages.go:94 info  发现工作区页面 名称:数学
2023-08-02 14:12:14 main.go:16 info  保存文件夹:/Users/melta/Documents/wolai/2023年08月02日14点12分
2023-08-02 14:12:14 main.go:54 info  正在导出密码学
2023-08-02 14:12:15 main.go:72 info  下载成功 文件名:密码学.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:15 main.go:54 info  正在导出其他
2023-08-02 14:12:19 main.go:72 info  下载成功 文件名:其他.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:19 main.go:54 info  正在导出哲学
2023-08-02 14:12:21 main.go:72 info  下载成功 文件名:哲学.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:21 main.go:54 info  正在导出政治
2023-08-02 14:12:22 main.go:72 info  下载成功 文件名:政治.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:22 main.go:54 info  正在导出英语
2023-08-02 14:12:24 main.go:72 info  下载成功 文件名:英语.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:24 main.go:44 info  忽略工作区页面 名称:计算机
2023-08-02 14:12:24 main.go:54 info  正在导出物理学
2023-08-02 14:12:25 main.go:72 info  下载成功 文件名:物理学.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:25 main.go:54 info  正在导出人生
2023-08-02 14:12:25 main.go:72 info  下载成功 文件名:人生.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:25 main.go:54 info  正在导出历史
2023-08-02 14:12:26 main.go:72 info  下载成功 文件名:历史.zip 链接:https://xxx.xxx.xxx/xxx.zip
2023-08-02 14:12:26 main.go:54 info  正在导出数学
2023-08-02 14:12:29 main.go:72 info  下载成功 文件名:数学.zip 链接:https://xxx.xxx.xxx/xxx.zip
```

可能的问题

1. 我的计算机页约有1k个文档，耗时了6分钟才完成，再多的情况是否会导致请求端口，我没有测试



## 其他说明

微信：SXL--LP（请备注Github哦）



## 项目编译

编译（版本：1.19.6）

```shell
go clone https://github.com/tengfei-xy/wolai.git
cd wolai
go build
```

