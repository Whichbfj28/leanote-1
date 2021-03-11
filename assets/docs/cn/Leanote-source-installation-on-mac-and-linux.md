## 本教程适合 Mac 及 linux 用户的**源码版**安装。

`leanote` 开发版(源码)适合开发人员，或需要快速更新功能的朋友。

-----------------------------
# 安装步骤:

1. 安装`Golang`。
2. 获取`Revel`和`Leanote`的源码。
3. 安装`Mongodb`。
4. 导入初始数据。
5. 使用`Revel`运行`Leanote`。


---------------------------------
## 1. 安装 `golang`

到 `golang.org` 官网下载最新版的`golang`, Leanote至少需要golang 1.7。如果被墙, 可以在 `http://golangtc.com/download` 下载。

以下为 1.8 版本的快速下载链接:

* linux 64位: http://www.golangtc.com/static/go/1.8/go1.8.darwin-amd64.pkg
* linux 32位: http://www.golangtc.com/static/go/1.8/go1.8.linux-386.tar.gz

假设将文件下载到 `/home/user1` 下, 解压文件：

```
$> cd /home/user1
$> tar -xzvf go1.6.linux-amd64.tar.gz
```

在 `/home/user1` 下新建一个目录`gopackage`, 这里面会放`go`的包和编译后的文件：

```
$> mkdir $GOPATH
```

配置环境变量, 编辑`/etc/profile`文件：
```
$> sudo vim /etc/profile
```

此处使用了`vim`文字编辑器，你可以使用自己喜欢的其他编辑器。在 `/etc/profile` 中添加以下几行：
```
export GOROOT=/home/user1/go
export GOPATH=$GOPATH
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

保存修改后，在终端运行以下命令使环境变量生效:
```
$> source /etc/profile
```

查看`go`是否安装成功:
```
$> go version
```
若出现类似以下信息证明安装成功
```
go version go1.6 linux/amd64
```

----------------------------
## 2. 获取`Revel`和 `Leanote` 的源码

### 2.1 方法1 （**推荐方法**）:

请下载 [leante-master.zip](https://github.com/coocn-cn/leanote/archive/master.zip)。解压后，将`src`文件夹复制到`$GOPATH/src/github.com/coocn-cn/leanote`

### 2.2 方法2

该方法使用`Golang`的 `go get` 来下载包, 这个命令会调用`git`, 所以必须先安装`git`。

- `ubuntu`下安装`git`:
```
$> sudo apt-get install git-core openssh-server openssh-client
```

- `centos`下安装`git`: 请参考: http://www.ccvita.com/370.html

获取`Leanote`:

打开终端, 以下命令会下载`Leanote`及依赖包, 时间可能会有点久, 请耐心等待。
```
$> go get github.com/coocn-cn/leanote
```

下载完成后，`Leanote`的源码在`$GOPATH/src/github.com/coocn-cn/leanote`下。

## 3. 安装`Mongodb`

### 3.1 安装`Mongodb`

到 [Mongodb 官网](http://www.mongodb.org/downloads) 下载相应系统的最新版安装包，或者从以下链接下载旧版本：

* 64位 linux Mongodb 3.0.1 下载链接(推荐): https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-3.0.1.tgz

下载到 `/home/user1`下, 直接解压即可:
```
$> cd /home/user1
$> tar -xzvf mongodb-linux-x86_64-3.0.1.tgz/
```

为了快速使用`mongodb`命令, 可以配置环境变量。编辑 `~/.profile`或`/etc/profile` 文件， 将`mongodb/bin`路径加入即可:
```
$> sudo vim /etc/profile
```
此处实例使用了`vim`文本编辑器，你可以使用自己熟悉的编辑器。

在`/etc/profile`中添加以下行，注意把用户名（`user1`）和相应的文件目录名（`mongodb-linux-x86_64-3.0.1`）替换成自己系统中的名称：
```
export PATH=$PATH:/home/user1/mongodb-linux-x86_64-3.0.1/bin
```

保存修改后，在终端运行以下命令使环境变量生效:
```
$> source /etc/profile
```


### 3.2 测试`Mongodb`安装

先在`/home/user1`下新建一个目录`data`存放`Mongodb`数据:
```
mkdir /home/user1/data
```

用以下命令启动`mongod`:
```
mongod --dbpath /home/user1/data
```

这时`mongod`已经启动，重新打开一个终端, 键入`mongo`进入交互程序：
```
$> mongo
> show dbs
...数据库列表
```

`Mongodb`安装到此为止, 下面为`Mongodb`导入`Leanote`初始数据。


-------------------------------------
## 4. 导入初始数据

`leanote` 初始数据在`$GOPATH/src/github.com/coocn-cn/leanote/mongodb_backup/leanote_install_data`中。

打开终端， 输入以下命令导入数据。

```
$> mongorestore -h localhost -d leanote --dir $GOPATH/src/github.com/coocn-cn/leanote/mongodb_backup/leanote_install_data
```

现在在`mongodb`中已经新建了`leanote`数据库, 可用命令查看下`Leanote`有多少张"表":
```
$> mongo
> show dbs #　查看数据库
leanote	0.203125GB
local	0.078125GB
> use leanote # 切换到leanote
switched to db leanote
> show collections # 查看表
files
has_share_notes
note_content_histories
note_contents
....
```

初始数据的`users`表中已有2个用户:
```
user1 username: admin, password: abc123 (管理员, 只有该用户才有权管理后台, 请及时修改密码)
user2 username: demo@leanote.com, password: demo@leanote.com (仅供体验使用)
```

-----------------------------------
## 5. 配置`Leanote`

`Leanote`的配置存储在文件 `conf/app.conf` 中。

请务必修改`app.secret`一项, 在若干个随机位置处，将字符修改成一个其他的值, 否则会有安全隐患!

其它的配置可暂时保持不变, 若需要配置数据库信息, 请参照 [Leanote问题汇总](https://github.com/coocn-cn/leanote/blob/master/assets/docs/cn/leanote-qa.md)。


---------------------------------
## 6. 编译`Leanote`

**注意:** 如果机器配置较差，那么编译可能需要较长时间。

新开一个窗口, 进入到目录`$GOPATH/src/github.com/coocn-cn/leanote`，运行:
```
$> ./assets/build/build.sh
```

命令执行成功输出`OK`后，编译出来的文件放在 `$GOPATH/src/github.com/coocn-cn/leanote/output/leanote` 目录下


---------------------------------
## 7. 运行`Leanote`

**注意:** 在此之前请确保`Mongodb`已在运行!

在编译的窗口继续运行:
```
$> ./output/leanote/bin/run.sh
```

等待程序输出 `Listening on.. 0.0.0.0:9000`
```
$> ./output/leanote/bin/run.sh
...
Listening on.. 0.0.0.0:9000
```

现在恭喜你, 打开浏览器输入: `http://localhost:9000` 体验Leanote吧!


-----------------------------------

# 注意!!!!!!!!!!!!!!
按照本教程启动`Mongodb`是没有权限控制的, 如果你的Leanote服务器暴露在外网, 任何人都可以访问你的Mongodb并修改, 所以这是极其危险的!!!!!!!!!!! 请务必为Mongodb添加用户名和密码并以`auth`启动, 方法请见: [为mongodb数据库添加用户](https://github.com/coocn-cn/leanote/wiki/QA#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%9F%9F%E5%90%8D)

# `leanote` 安装/配置问题汇总

如果运行有问题或想要进一步配置`Leanote`, 请参照 [Leanote问题汇总](https://github.com/coocn-cn/leanote/wiki/QA)。
