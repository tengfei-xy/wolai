# wolai
wolai的一些便捷功能，支持导出所有文章（怕官方跑路哈哈哈哈）



## 使用方法

[下载地址](https://github.com/tengfei-xy/wolai/releases)

1. 首次运行将生成config.yaml

2. 修改配置文件后，重新运行主程序



## 配置文件说明

```yaml
# 请从浏览器登录自己的个人空间后获取
cookie: "isg=xxx; token=xxx, wolai_client_id=xxx"

# 作为保存的目标文件夹
# 建议为绝对路径
backupBackupDir: /Users/melta/Documents/wolai

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



## 系统说明

1. windows的配置文件需要保存为ANSI编码，其他系统使用utf-8



## 主要功能-导出全部md

> 注：由于官方目前没有给出导出md的正式[API](https://www.wolai.com/wolai/7FB9PLeqZ1ni9FfD11WuUi)，所以是直接用了浏览器与后端直接交互接口

使用方法

```
./wolai
```

正常日志

```
XiaoXiaoFeiTongXue  wolai git:(main)   go run .                 
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:哲学
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:数学
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:政治
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:法律
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:英语
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:历史
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:计算机
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:密码学
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:物理学
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:世界 页面:心理学
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:生活
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:思考
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:价值观
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:人生观
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:世界观
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:飞常人生 页面:其他
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:MySQL
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:Python
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:任务与理解
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:奔赴山海
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:所谓伊人
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:执子之手
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:路过人间
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:纸短情长
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:linux
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:VPS
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:产品
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:用户
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:网络
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:公司
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:机房
2023-11-23 10:35:53 workspace.go:263 info  发现工作区:公司 页面:AD域
2023-11-23 10:35:53 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界
2023-11-23 10:35:53 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/飞常人生
2023-11-23 10:35:53 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/数据分析
2023-11-23 10:35:53 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er
2023-11-23 10:35:53 main.go:26 info  创建 保存目标文件夹:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司
2023-11-23 10:35:53 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:哲学
2023-11-23 10:35:57 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/哲学.zip 链接:xxxxx
2023-11-23 10:35:57 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:数学
2023-11-23 10:36:04 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/数学.zip 链接:xxxxx
2023-11-23 10:36:04 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:政治
2023-11-23 10:36:05 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/政治.zip 链接:xxxxx
2023-11-23 10:36:05 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:法律
2023-11-23 10:36:07 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/法律.zip 链接:xxxxx
2023-11-23 10:36:07 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:英语
2023-11-23 10:36:09 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/英语.zip 链接:xxxxx
2023-11-23 10:36:09 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:历史
2023-11-23 10:36:10 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/历史.zip 链接:xxxxx
2023-11-23 10:36:10 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:计算机
2023-11-23 10:38:18 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/计算机.zip 链接:xxxxx
2023-11-23 10:38:18 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:密码学
2023-11-23 10:38:19 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/密码学.zip 链接:xxxxx
2023-11-23 10:38:19 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:物理学
2023-11-23 10:38:20 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/物理学.zip 链接:xxxxx
2023-11-23 10:38:20 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:世界 页面:心理学
2023-11-23 10:38:21 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/世界/心理学.zip 链接:xxxxx
2023-11-23 10:38:21 export.go:37 warn  忽略导出 工作区:梦娜·学不进·想自莎 子空间:飞常人生
2023-11-23 10:38:21 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:MySQL
2023-11-23 10:38:22 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/数据分析/MySQL.zip 链接:xxxxx
2023-11-23 10:38:22 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:Python
2023-11-23 10:38:24 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/数据分析/Python.zip 链接:xxxxx
2023-11-23 10:38:24 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:数据分析 页面:任务与理解
2023-11-23 10:38:25 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/数据分析/任务与理解.zip 链接:xxxxx
2023-11-23 10:38:25 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:奔赴山海
2023-11-23 10:39:00 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er/奔赴山海.zip 链接:xxxxx
2023-11-23 10:39:00 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:所谓伊人
2023-11-23 10:39:01 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er/所谓伊人.md 链接:xxxxx
2023-11-23 10:39:01 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:执子之手
2023-11-23 10:39:01 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er/执子之手.md 链接:xxxxx
2023-11-23 10:39:01 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:路过人间
2023-11-23 10:39:02 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er/路过人间.md 链接:xxxxx
2023-11-23 10:39:02 export.go:27 info  开始导出 工作区:梦娜·学不进·想自莎 子空间:我们的那些事er 页面:纸短情长
2023-11-23 10:39:26 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/梦娜·学不进·想自莎/我们的那些事er/纸短情长.zip 链接:xxxxx
2023-11-23 10:39:26 export.go:25 info  开始导出 工作区:公司 页面:linux
2023-11-23 10:39:27 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/linux.zip 链接:xxxxx
2023-11-23 10:39:27 export.go:25 info  开始导出 工作区:公司 页面:VPS
2023-11-23 10:39:28 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/VPS.zip 链接:https
2023-11-23 10:39:28 export.go:25 info  开始导出 工作区:公司 页面:产品
2023-11-23 10:39:29 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/产品.zip 链接:xxxxx
2023-11-23 10:39:29 export.go:25 info  开始导出 工作区:公司 页面:用户
2023-11-23 10:39:31 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/用户.zip 链接:xxxxx
2023-11-23 10:39:31 export.go:25 info  开始导出 工作区:公司 页面:网络
2023-11-23 10:39:31 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/网络.zip 链接:xxxxx
2023-11-23 10:39:31 export.go:18 warn  忽略导出 工作区:公司 页面:公司
2023-11-23 10:39:31 export.go:25 info  开始导出 工作区:公司 页面:机房
2023-11-23 10:39:33 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/机房.zip 链接:xxxxx
2023-11-23 10:39:33 export.go:25 info  开始导出 工作区:公司 页面:AD域
2023-11-23 10:39:34 export.go:90 info  下载成功 保存路径:/Users/melta/Documents/wolai/2023年11月23日10点35分/公司/AD域.zip 链接:xxxxx
2023-11-23 10:39:34 main.go:104 info  导出结束!欢迎再次使用
XiaoXiaoFeiTongXue  wolai git:(main)   
```

可能的问题

1. 我的计算机页约有1k个文档，耗时了2分钟完成，更多的md的情况是否会导致请求端口，我没有测试



## 其他说明

微信：SXL--LP（请备注Github哦）



## 项目编译

编译（版本：1.19.6）

```shell
go clone https://github.com/tengfei-xy/wolai.git
cd wolai
go build
```

