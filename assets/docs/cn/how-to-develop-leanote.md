## 安装Leanote

* [`Leanote` 源码导读](https://github.com/coocn-cn/leanote/blob/master/assets/docs/cn/how-to-read-leanote-source.md)
* [Leanote开发版安装教程](https://github.com/coocn-cn/leanote/blob/master/assets/docs/cn/Leanote开发版在Cubieboard上详细安装教程.md)

## 运行Leanote
以开发模式运行Leanote:
```
$> revel run github.com/leanote/leanote dev [port]
如:
$> revel run github.com/leanote/leanote dev 9000
```

以生产模式运行Leanote:

```
$> revel run github.com/leanote/leanote prod [port]
如:
$> revel run github.com/leanote/leanote prod 8080
```

注意, 以开发模式运行Leanote笔记主页使用的view是note-dev.html, 以生产模式运行Leanote笔记主页使用的view是note.html.

## 使用eclipse开发Leanote

首先确保你的eclipse已安装`goclipse`插件, 然后直接将leanote导入到eclipse中(leanote已是一个eclipse项目, 直接导入即可)

为了确保eclipse能自动编译leanote, 你可以建一个symlink src 指向app(src是goclipse默认的源码路径)

当然, 你还可以用Sublime开发Leanote.

## 构建 Leanote 静态文件
以开发模式启动Leanote时, 使用的是note-dev.html, 发布Leanote时需要合并js, note-dev.html转成note.html.

Leanote前端通过gulp编译, 请在leanote根目录下执行:
```
$> npm install
$> gulp
```
在此之前你可能需要安装nodejs, gulp.