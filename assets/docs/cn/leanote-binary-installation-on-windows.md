## 本教程适合 Windows 用户的**二进制版**安装。

- Windows 用户的源码版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E6%BA%90%E7%A0%81%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Windows)。
- Mac, Linux 用户的二进制版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Mac-and-Linux)。
- Mac, Linux 用户的源码版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E6%BA%90%E7%A0%81%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Mac-and-Linux)。

----------------------
# 安装步骤:

1. 下载 `leanote` 二进制版。
2. 安装 `mongodb`。
3. 导入初始数据。
4. 配置 `leanote`。
5. 运行 `leanote`。


-------------------------------
## 1. 下载 `leanote` 二进制版

下载 [leanote 最新二进制版](http://leanote.org/#download), 请根据系统选择相应文件。

假设将文件下载到 `C:\user1` 下并解压, 现在应该有 `C:\users1\leanote`。

----------------------------
## 2. 安装 `mongodb`

## 2.1 安装 `mongodb`

到 [mongodb 官网](http://www.mongodb.org/downloads) 下载相应系统的最新版安装包。一直点击下一步默认安装。采用默认设置或自定义安装，如下图所示：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image006.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image008.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image009.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image010.png)

点击 **Finish** 安装完毕!

### 2.2 测试`mongodb`安装

在`C`盘根目录下建立`dbanote`目录用于放置笔记的数据文件：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image011.png)

直接按`WinKey+R`, 输入`cmd` 打开命令行，输入以下命令（不含`C:\>`）：
```
C:\>mongod --dbpath C:\dbanote 
```

启动数据库，界面如下：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image012.png)

####**★注意:此时这个命令行窗口最小化，不要关闭！切记！！！** 

重新打开一个终端 (直接按`WinKey+R`, 输入`cmd` 打开命令行），输入:
```
C:\> mongo
```
行首出现`>` 表示进入`mongo` 的交互程序。此时输入：
```
> show dbs
```

如下图：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image014.png)

`mongodb` 到此安装完成！

-----------------------------------------

## 3. 导入初始数据

按`win+R`，输入`cmd`，回车，打开新的命令行，复制并运行以下命令。注意对应你安装的`mongdb`的版本：

- `mongodb v2` 的导入命令为:
```
mongorestore  -h localhost -d leanote  --directoryperdb C:\user1\leanote\mongodb_backup\leanote_install_data
```

- `mongodb v3` 的导入命令为:
```
mongorestore -h localhost -d leanote --dir C:\user1\leanote\mongodb_backup\leanote_install_data
```

完成数据导入，如下图：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image017.png)

为测试导入数据，继续在导入数据的命令行输入：
```
C:\> mongo
> show dbs          # 查看数据库
admin    (empty)
leanote  0.078GB        # Leanote 导入成功的数据库
local    0.078GB 
```

**注意：导入成功的数据已经包含2个用户**
```
user1 username: admin, password: abc123 (管理员, 只有该用户可以管理后台)  
user2 username: demo@leanote.com, password: demo@leanote.com (仅供体验使用)
```

----------------------------------
## 4. 配置`leanote`

`leanote`的配置存储在文件 `conf/app.conf` 中。

请务必修改`app.secret`一项, 在若干个随机位置处，将字符修改成一个其他的值, 否则会有安全隐患!

其它的配置可暂时保持不变, 若需要配置数据库信息, 请参照 [leanote问题汇总](https://github.com/leanote/leanote/wiki/QA)。

---------------------------------
## 5. 运行`leanote`

以 **管理员权限** 打开`cmd`，输入：

```
$> cd C:\users\leanote\bin
$> run.bat
```

最后出现以下信息证明运行成功:
```
...
TRACE 2013/06/06 15:01:27 watcher.go:72: Watching: /home/life/leanote/bin/src/github.com/leanote/leanote/conf/routes
Go to /@tests to run the tests.
Listening on :9000...
```

恭喜你, 打开浏览器输入: `http://localhost:9000` 体验`leanote`吧!

# 注意!!!!!!!!!!!!!!
按照本教程启动`Mongodb`是没有权限控制的, 如果你的Leanote服务器暴露在外网, 任何人都可以访问你的Mongodb并修改, 所以这是极其危险的!!!!!!!!!!! 请务必为Mongodb添加用户名和密码并以`auth`启动, 方法请见: [为mongodb数据库添加用户](https://github.com/leanote/leanote/wiki/QA#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%9F%9F%E5%90%8D)

# `leanote` 安装/配置问题汇总

如果运行有问题或想要进一步配置`leanote`, 请参照 [leanote问题汇总](https://github.com/leanote/leanote/wiki/QA)。
