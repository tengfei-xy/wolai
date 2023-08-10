# wolai
wolai的一些便捷功能，支持导出所有文章（怕官方跑路哈哈哈哈）



## 使用方法

[下载地址](https://github.com/tengfei-xy/wolai/releases)

1.首次运行将生成config.yaml.tmp

2.修改配置文件后并重命名为config.yaml

3.重新运行主程序



## 配置文件说明

```yaml
# 请从浏览器登录自己的个人空间后获取
login:
    cookie: "isg=xxx; token=xxx, wolai_client_id=xxx"
    
# 作为保存的目标文件夹
save:
    targetpath: "/Users/melta/Documents/wolai"

# 忽略的页面名称
ignore:
    - spaceName: "空间名"
      pageName:
        - ""
        - ""
    - spaceName: "空间名"
      pageName:
      # 表示忽略所有
        - "*"

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
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:人生
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:历史
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:政治
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:英语
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:计算机
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:密码学
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:梦娜·学不进·想自莎 页面名称:物理学
2023-08-10 12:36:08 pages.go:41 info  发现 工作区名称:233 页面名称:233
2023-08-10 12:36:08 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年08月10日12点36分
2023-08-10 12:36:08 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:人生
2023-08-10 12:36:10 export.go:76 info  下载成功 文件名:人生.zip 链接:https://x.x.x/x.zip
2023-08-10 12:36:10 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:历史
2023-08-10 12:36:11 export.go:76 info  下载成功 文件名:历史.zip 链接:https://x.x.x/x.zip
2023-08-10 12:36:11 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:政治
2023-08-10 12:36:12 export.go:76 info  下载成功 文件名:政治.zip 链接:https://x.x.x/x.zip
2023-08-10 12:36:12 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:英语
2023-08-10 12:36:13 export.go:76 info  下载成功 文件名:英语.zip 链接:https://x.x.x/x.zip
2023-08-10 12:36:13 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:计算机
2023-08-10 12:37:20 export.go:76 info  下载成功 文件名:计算机.zip 链接:https://x.x.x/x.zip
2023-08-10 12:37:20 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:密码学
2023-08-10 12:37:20 export.go:76 info  下载成功 文件名:密码学.zip 链接:https://x.x.x/x.zip
2023-08-10 12:37:20 export.go:59 info  导出 空间名:梦娜·学不进·想自莎 页面名:物理学
2023-08-10 12:37:21 export.go:76 info  下载成功 文件名:物理学.zip 链接:https://x.x.x/x.zip
2023-08-10 12:37:21 export.go:59 info  导出 空间名:233 页面名:233
2023-08-10 12:37:24 export.go:76 info  下载成功 文件名:233.zip 链接:https://x.x.x/x.zip
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

