## 本教程适合 Windows 用户的**源码版**安装。

- Windows 用户的二进制版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Windows)。
- Mac, Linux 用户的源码版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E6%BA%90%E7%A0%81%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Mac-and-Linux)。
- Mac, Linux 用户的二进制版安装，参见[这里](https://github.com/leanote/leanote/wiki/Leanote-%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%89%88%E8%AF%A6%E7%BB%86%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B----Mac-and-Linux)。

注意：为增加本程序兼容性，请尽量按照本程序操作（ **`32位系统`**，源码安装位置：**`C盘`** ），如需要自定义环境，请随机应变。

-----------------------------
# 安装步骤:

1. 下载安装文件。
2. 安装`golang`。
3. 安装`mongodb`。
4. 安装`leanote`源码。
5. 导入初始数据。
6. 配置`leanote`。
7. 运行`leanote`。


------------------

## 1. 下载安装文件（以32位为例）

* `golang`(1.7+)下载: http://www.golangtc.com/static/go/1.8/go1.8.windows-386.msi
* `mongodb`下载: https://fastdl.mongodb.org/win32/mongodb-win32-i386-2.6.8-signed.msi?_ga=1.163324924.1783433278.1426342651  
* `leanote-all` 依赖环境与源码下载: https://github.com/leanote/leanote-all/archive/master.zip  

------------------------
## 2. 安装`golang`

使用下载的`golang`安装包安装：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image001.png)

一直点击点击下一步，默认安装。
安装完成后，直接按`WinKey+R`， 输入`cmd` 打开命令行，输入 ```go version``` 如出现如下显示，说明`golang`安装正确。

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image003.png)

给增加`GO`添加 `GOPATH` 和 `GOROOT` 环境变量:
右键我的电脑 — 属性 – 高级 – 环境变量 – 如下图

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image005.png)

注意俩个变量的区别。

-------------------------------------
## 3. 安装`mongodb`

### 3.1 安装 `mongodb`

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image006.png)

与Golang一样一直点击下一步默认安装。如需自定义设置，在第二部如下图选择：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image008.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image009.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image010.png)

点击 **Finish** 安装完毕!

### 3.2 测试`mongodb`

在`C`盘根目录下建立`dbanote`目录用于放置笔记的数据文件，如下图：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image011.png)

直接按`WinKey+R`， 输入`cmd` 打开命令行，输入以下命令（不含`C:\>`）:
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

--------------------------------------
## 4. 安装`leanote`源码

解压之前下载的 `leanote-all` 源码包，将`src`文件夹复制或移动到 `C:\Go` 下（如下图示），如出现覆盖确认对话框，点击确认。

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image016.png)

-----------------------
## 5. 导入初始数据

按`win+R`，输入`cmd`，回车，打开新的命令行，复制并运行以下命令。注意对应你安装的`mongdb`的版本：

- `mongodb v2` 的导入命令为:
```
mongorestore  -h localhost -d leanote  --directoryperdb C:\Go\src\github.com\leanote\leanote\mongodb_backup\leanote_install_data
```

- `mongodb v3` 的导入命令为:
```
mongorestore -h localhost -d leanote --dir C:\Go\src\github.com\leanote\leanote\mongodb_backup\leanote_install_data
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
user2 username: demo@leanote.com, password: demo@leanote.com (仅共体验使用)
```

---------------------
## 6. 配置`leanote`

`leanote`的配置存储在文件 `conf/app.conf` 中。

请务必修改`app.secret`一项, 在若干个随机位置处，将字符修改成一个其他的值, 否则会有安全隐患! 

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image019.png)

其它的配置可暂时保持不变, 若需要配置数据库信息, 请参照 [leanote问题汇总](https://github.com/leanote/leanote/wiki/QA)。

---------------------------------
## 7. 运行`leanote`

在新打开的命令行窗口输入：
```
go install github.com\revel\cmd\revel
```
生成`revel`命令，继续输入：
```
revel run github.com\\leanote\\leanote  
```

正常启动界面 如下图：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image021.png)

####**★注意：此时这个命令行窗口不关闭，最小化（与之前的`mongodb`命令行窗口一样）**

到此，Windows下安装`leanote`正式结束，记得之前的  **俩个命令行窗口不能关闭**。
现在你就可以打开浏览器，输入`http://localhost`，用之前导入原始数据包含的用户:

```
user1 username: admin, password: abc123 (管理员, 只有该用户可以管理后台)
user2 username: demo@leanote.com, password: demo@leanote.com (仅共体验使用)
```

来访问你的自建笔记环境了。

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image023.png)


-----------------------------------

# 注意!!!!!!!!!!!!!!
按照本教程启动`Mongodb`是没有权限控制的, 如果你的Leanote服务器暴露在外网, 任何人都可以访问你的Mongodb并修改, 所以这是极其危险的!!!!!!!!!!! 请务必为Mongodb添加用户名和密码并以`auth`启动, 方法请见: [为mongodb数据库添加用户](https://github.com/leanote/leanote/wiki/QA#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%9F%9F%E5%90%8D)

# `leanote` 安装/配置问题汇总

如果运行有问题或想要进一步配置`leanote`, 请参照 [leanote问题汇总](https://github.com/leanote/leanote/wiki/QA)。

--------------------------------

##**问题汇总**

### 问题 0

"no reachable server" 

请修改conf/app.conf中的 db.host=localhost 为 db.host=127.0.0.1 再重启leanote

### **问题1：**

```
Go to /@tests to run the tests.
panic: auth fails

goroutine 1 [running]:
github.com/leanote/leanote/app/db.Init()
/home/life/gopackage1/src/github.com/leanote/leanote/app/db/Mgo.go:64 +0x356  
```
解答:

数据库配置有问题, 请修改`leanote/conf/app.conf`文件, 是否用户名和密码配置有误?


### **问题2: 修改默认80端口?**

修改`leanote/conf/app.conf`, 比如改成9000
```
http.port=9000

site.url=http://localhost:9000
```

### **问题3: 为数据库添加用户**

建立数据库用户：
打开命令行窗口输入：
```
C:\ >mongo
# 首先切换到leanote数据库下
> use leanote;
# 添加一个用户root, 密码是abc123
> db.addUser("root", "abc123");
{
    "_id" : ObjectId("53688d1950cc1813efb9564c"),
    "user" : "root",
    "readOnly" : false,
    "pwd" : "e014bfea4a9c3c27ab34e50bd1ef0955"
}
# 测试下是否正确
> db.auth("root", "abc123");
1   # 返回1表示正确
```

如下图：
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image025.png)


用户添加好后重新运行下mongodb, 并开启权限验证. 在mongod的终端按`ctrl+c`即可退出mongodb.

```
# 重新启动mongodb:

$> mongod --dbpath C:\Dbanote  –auth  
```
如下图：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image027.png)

其它的配置请保持不变, 若需要配置数据库信息, 请查看下文"问题3"
修改``C:\Go\src\github.com\leanote\leanote\confc\app.conf``, mongodb的配置一般只需要修改`db.username`和`db.password`就行了
如下图：（强烈建议使用**Notepad++**类编辑器修改）

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image029.png)