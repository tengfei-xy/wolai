# wolai
wolai的一些便捷功能，支持导出所有文章（怕官方跑路哈哈哈哈）



## 项目编译

编译（版本：1.19.6）

```
go build
```



## 使用前提

需要在源码中填充cookie变量（包含isg、token、wolai_client_id值）



## 主要功能-导出md

> 注：由于官方目前没有给出导出md的正式[API](https://www.wolai.com/wolai/7FB9PLeqZ1ni9FfD11WuUi)，所以是直接用了浏览器与后端直接交互接口

使用方法

```
./wolai
```

日志

```
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:人生
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:哲学
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:历史
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:政治
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:密码学
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:物理学
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:数学
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:英语
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:计算机
2023-08-01 17:05:34 pages.go:94 info  获取ID成功:其他
2023-08-01 17:05:36 main.go:40 info 下载成功 文件:人生.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:39 main.go:40 info 下载成功 文件:哲学.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:40 main.go:40 info 下载成功 文件:历史.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:41 main.go:40 info 下载成功 文件:政治.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:42 main.go:40 info 下载成功 文件:密码学.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:42 main.go:40 info 下载成功 文件:物理学.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:47 main.go:40 info 下载成功 文件:数学.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:05:49 main.go:40 info 下载成功 文件:英语.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:11:44 main.go:40 info 下载成功 文件:计算机.zip 链接:https://xxx.xxx.xxx/xxx.zip  
2023-08-01 17:11:48 main.go:40 info 下载成功 文件:其他.zip 链接:https://xxx.xxx.xxx/xxx.zip 
```

可能的问题

1. 我的计算机页约有1k个文档，耗时了6分钟才完成，再多的情况是否会导致请求端口，我没有测试



## 其他说明

微信（请备注Github哦）：SXL--LP

