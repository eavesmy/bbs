<p align = "center">
<img alt="Gorobbs" src="#">
<br><br>
轻而快，为未来而构建
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/b3log/pipe"><img src="https://img.shields.io/travis/b3log/pipe.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/b3log/pipe"><img src="https://goreportcard.com/badge/github.com/b3log/pipe?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/b3log/pipe/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/b3log/pipe.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/b3log/pipe"><img src="https://img.shields.io/github/languages/code-size/b3log/pipe.svg?style=flat-square"></a>
<a title="GPLv3" target="_blank" href="https://github.com/b3log/pipe/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPLv3-orange.svg?style=flat-square"></a>
<br>
<a title="Releases" target="_blank" href="https://github.com/b3log/pipe/releases"><img src="https://img.shields.io/github/release/b3log/pipe.svg?style=flat-square"></a>
<a title="Release Date" target="_blank" href="https://github.com/b3log/pipe/releases"><img src="https://img.shields.io/github/release-date/b3log/pipe.svg?style=flat-square&color=99CCFF"></a>
<a title="Docker Pulls" target="_blank" href="https://hub.docker.com/r/b3log/pipe"><img src="https://img.shields.io/docker/pulls/b3log/pipe.svg?style=flat-square&color=blueviolet"></a>
<a title="Docker Image Size" target="_blank" href="https://hub.docker.com/r/b3log/pipe"><img src="https://img.shields.io/microbadger/image-size/b3log/pipe.svg?style=flat-square&color=ff96b4"></a>
<br>
<a title="GitHub Commits" target="_blank" href="https://github.com/b3log/pipe/commits/master"><img src="https://img.shields.io/github/commit-activity/m/b3log/pipe.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/b3log/pipe/commits/master"><img src="https://img.shields.io/github/last-commit/b3log/pipe.svg?style=flat-square&color=FF9900"></a>
<a title="GitHub Pull Requests" target="_blank" href="https://github.com/b3log/pipe/pulls"><img src="https://img.shields.io/github/issues-pr-closed/b3log/pipe.svg?style=flat-square&color=FF9966"></a>
<a title="Hits" target="_blank" href="https://github.com/b3log/hits"><img src="https://hits.b3log.org/b3log/pipe.svg"></a>
<br><br>
<a title="GitHub Watchers" target="_blank" href="https://github.com/b3log/pipe/watchers"><img src="https://img.shields.io/github/watchers/b3log/pipe.svg?label=Watchers&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Stars" target="_blank" href="https://github.com/b3log/pipe/stargazers"><img src="https://img.shields.io/github/stars/b3log/pipe.svg?label=Stars&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Forks" target="_blank" href="https://github.com/b3log/pipe/network/members"><img src="https://img.shields.io/github/forks/b3log/pipe.svg?label=Forks&style=social"></a>&nbsp;&nbsp;
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>
</p>

## 简介

[Gorobbs](https://github.com/letseeqiji/gorobbs) 是一款小而轻的开源BBS，，专为普通用户设计。我们的目标是打造最轻量化的分布式BBS系统！

## 案例

* [新书来了](http://www.xinshulaile.com)


## 功能

* 多用户BBS
* 自定义导航
* 多主题 / 多语言
* 可配置动静分离
* 支持 MySQL

## 界面

### 开始使用

![start](https://#)

### 管理后台

![console](https://)

### 编辑文章

![post](https://)


## 安装

### 本地试用

* [下载](https://github.com/b3log/pipe/releases)最新的发布包解压，进入解压目录运行 pipe/pipe.exe
* 从源码构建可参考[这里](https://hacpai.com/article/1533965022328)

**请注意**：我们不建议通过发布包或者源码构建部署，因为这样的部署方式在将来有新版本发布时升级会比较麻烦。
这两种方式请仅用于本地试用，线上生产环境建议通过 Docker 部署。

### Docker 部署

获取最新镜像：

```shell
docker pull letseeqiji/gorobbs
```

* 使用 MySQL

  先手动建库（库名 `gorobbs`，字符集使用 `utf8mb4`，排序规则 `utf8mb4_general_ci`），然后启动容器：
  
  ```shell
  docker run --detach --name gorobbs --network=host \
      letseeqiji/gorobbs --mysql="root:123456@(127.0.0.1:3306)/gorobbs?charset=utf8mb4&parseTime=True&loc=Local" --runtime_mode=prod --port=5897 --server=http://localhost:5897
  ```
  为了简单，使用了主机网络模式来连接主机上的 MySQL。
  

  
启动参数说明：

* `--port`：进程监听端口
* `--server`：访问时的链接

完整启动参数的说明可以使用 `-h` 来查看。

### Docker 升级

1. 拉取最新镜像
2. 重启容器

可参考[这里](https://#)编写一个重启脚本，并通过 crontab 每日凌晨运行来实现自动更新。

## 文档

* [《提问的智慧》精读注解版](https://#)
* [用户指南](https://#)
* [开发指南](https://#)
* [主题开发指南](https://#)
* [贡献指南](https://#)

## 社区

* [讨论区](https://#)
* [报告问题](https://#)

## 授权

Gorobbs 使用 [MIT](https://#) 开源协议。

## 鸣谢

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库


---

## 开源项目推荐

