效果：

这个导航页前端源于：

- https://github.com/TopVitamin/static-nav

直接使用它有一个弊端就是我如果要添加或者修改链接项是需要去编辑 HTML 的，所以我要给它加个后端优化一下。

但是它本身也就是做一个静态页面导航用，我就不想把它做的太复杂，所以也没有引入数据库啥的，页面的链接信息我就直接把它保存到
Yaml 中，如果我们要修改页面的链接就直接去编辑这个 Yaml 文件即可，格式如下：

```yaml
title: 导航
favicon: dp2.png
mainTitle: 嗖嗖嗖
inputPlaceholder: 不懂么？不懂就搜，给老子搜！
groups:
  - group: test
    items:
      - name: 百度
        url: https://www.baidu.com
      - name: 必应
        url: https://www.biying.com
  - group: test2
    items:
      - name: 必应
        url: https://www.biying.com
      - name: 百度
        url: https://www.baidu.com
```

我们只需要在二进制程序所在目录的 `conf` 目录下放一个内容格式如上的 `data.yaml` 文件，直接运行程序即可享用，如下：

```bash
$ ./NavIndex 
Iris Version: 12.2.0

Now listening on: http://0.0.0.0:8080
Application started. Press CTRL+C to shut down.
```

二进制程序可以直接在 Releases 区下载，暂时只提供了 Linux 版本。

或者你可以直接使用 Docker 一键运行：

```bash
# - 实际部署时可以通过 -v <宿主机目录>:/opt/conf 将 data.yaml 所在目录映射出来
# - 如果在 kubernetes 中运行的话 data.yaml 可以保存到 ConfigMap 然后挂载到 /opt/conf 目录
docker run --name nav-index -p8080:8080 zze326/navindex:v1.0
```
