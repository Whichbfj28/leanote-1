本文是关于Leanote安装/配置的一些问题 ([English version](https://github.com/leanote/leanote/wiki/Leanote-QA-English)), 关于Leanote的使用教程请参照 [Leanote使用教程集](http://leanote.leanote.com/post/Leanote-manual-project).

在使用或安装Leanote的任何疑问, 请加Leanote官方QQ群 256076853 进行讨论 或 联系leanote@leanote.com.

欢迎为Leanote作贡献, 一个issue或在文档上一些改进都是对Leanote最好的支持!

目录:

* [运行出错 "no reachable server"](https://github.com/leanote/leanote/wiki/QA#no-reachable-server)
* [安装Leanote后运行出错](https://github.com/leanote/leanote/wiki/QA#%E5%AE%89%E8%A3%85leanote%E5%90%8E%E8%BF%90%E8%A1%8C%E5%87%BA%E9%94%99)
* [Leanote运行成功, 但不能登录](https://github.com/leanote/leanote/wiki/QA#leanote%E8%BF%90%E8%A1%8C%E6%88%90%E5%8A%9F-%E4%BD%86%E4%B8%8D%E8%83%BD%E7%99%BB%E5%BD%95)
* [修改Leanote运行端口](https://github.com/leanote/leanote/wiki/QA#%E4%BF%AE%E6%94%B9leanote%E8%BF%90%E8%A1%8C%E7%AB%AF%E5%8F%A3)
* [如何绑定域名?](https://github.com/leanote/leanote/wiki/QA#%E5%A6%82%E4%BD%95%E7%BB%91%E5%AE%9A%E5%9F%9F%E5%90%8D)
* [为mongodb数据库添加用户](https://github.com/leanote/leanote/wiki/QA#%E4%B8%BAmongodb%E6%95%B0%E6%8D%AE%E5%BA%93%E6%B7%BB%E5%8A%A0%E7%94%A8%E6%88%B7)
* [为Leanote指定超级管理员帐户(admin用户)](https://github.com/leanote/leanote/wiki/QA#%E4%B8%BAleanote%E6%8C%87%E5%AE%9A%E8%B6%85%E7%BA%A7%E7%AE%A1%E7%90%86%E5%91%98%E5%B8%90%E6%88%B7admin%E7%94%A8%E6%88%B7)
* [为Leanote配置https](https://github.com/leanote/leanote/wiki/QA#%E4%B8%BAleanote%E9%85%8D%E7%BD%AEhttps)
* [Import of github.com/revel/revel/modules/testrunner failed](https://github.com/leanote/leanote/wiki/QA#import-of-githubcomrevelrevelmodulestestrunner-failed)
* [开发版如何更新leanote?](https://github.com/leanote/leanote/wiki/QA#%E5%BC%80%E5%8F%91%E7%89%88%E5%A6%82%E4%BD%95%E6%9B%B4%E6%96%B0leanote)
* [二进制版如何更新leanote?](https://github.com/leanote/leanote/wiki/QA#%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%89%88%E5%A6%82%E4%BD%95%E6%9B%B4%E6%96%B0leanote)
* [为什么需要 site.url](https://github.com/leanote/leanote/wiki/QA#%E4%B8%BA%E4%BB%80%E4%B9%88%E9%9C%80%E8%A6%81-siteurl)
* [客户端不能同步图片](https://github.com/leanote/leanote/wiki/QA#%E5%AE%A2%E6%88%B7%E7%AB%AF%E4%B8%8D%E8%83%BD%E5%90%8C%E6%AD%A5%E5%9B%BE%E7%89%87)
* [导出PDF配置 wkhtmltopdf](https://github.com/leanote/leanote/wiki/QA#%E5%AF%BC%E5%87%BApdf%E9%85%8D%E7%BD%AE-wkhtmltopdf)
* [不能通过IP访问](https://github.com/leanote/leanote/wiki/QA#%E4%B8%8D%E8%83%BD%E9%80%9A%E8%BF%87ip%E8%AE%BF%E9%97%AE)
## no reachable server

请确保数据库是否启动, 如果确定已启动 可以 尝试将 conf/app.conf `db.host=localhost` 改为 `db.host=127.0.0.1`

修改后请重新启动Leanote.

## 安装Leanote后运行出错

如果出现以下问题:

```
Go to /@tests to run the tests.
panic: auth fails

goroutine 1 [running]:
github.com/leanote/leanote/app/db.Init()
	/home/life/gopackage1/src/github.com/leanote/leanote/app/db/Mgo.go:64 +0x356
```

Leanote运行出错90%的原因是数据库的问题, 请检查:

1. 数据库是否启动了?
2. 如果数据库是以auth方式启动的, 请检查conf/app.conf是否配置了正确的数据库用户名和密码

注意, 默认的conf/app.conf 数据库配置如下:
```
# mongdb
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username= # if not exists, please leave it blank
db.password= # if not exists, please leave it blank
# or you can set the mongdb url for more complex needs the format is:
# mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
# db.url=mongodb://root:root123@localhost:27017/leanote
db.urlEnv=${MONGODB_URL} # set url from env
```
即数据库名为leanote, 用户名和密码为空, 请检查是否正确.

## Leanote运行成功, 但不能登录

原因: 数据库已停止运行, 请重新启动数据库和Leanote.

如果数据库在运行, 请重新启动Leanote.

## 修改Leanote运行端口

比如想以8080端口启动.

修改conf/app.conf:
```
http.port=8080
site.url=http://localhost:8080
```

请重启Leanote, 使用http://localhost:8080访问.

## 如何绑定域名?
比如想绑定域名a.com到你运行Leanote服务器, 你需要将leanote以80端口运行, 请修改conf/app.conf的如下配置:
```
http.port=80
site.url=http://a.com
```
然后启动Leanote.
当然你还需要将a.com绑定ip到Leanote服务器.

如果服务器上已有其它程序运行了80端口, 怎么办呢? 请google或百度下 "使用nginx分发请求到不同端口".

## 为mongodb数据库添加用户

像mysql一样有root用户, mongodb初始是没有用户的, 这样很不安全, 所以要为leanote数据库新建一个用户来连接leanote数据库(注意, 并不是为leanote的表users里新建用户, 而是新建一个连接leanote数据库的用户, 类似mysql的root用户).

mognodb v2与v3创建用户命令有所不同

mongodb v2 创建用户如下:
```
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
1 # 返回1表示正确
```
mongodb v3 创建用户如下:
```
# 首先切换到leanote数据库下
> use leanote;
# 添加一个用户root, 密码是abc123
> db.createUser({
    user: 'root',
    pwd: 'abc123',
    roles: [{role: 'dbOwner', db: 'leanote'}]
});
# 测试下是否正确
> db.auth("root", "abc123");
1 # 返回1表示正确
```

用户添加好后重新运行下mongodb, 并开启权限验证. 在mongod的终端按ctrl+c即可退出mongodb.

启动mongodb:
```
$> mongod --dbpath /home/user1/data --auth
```

**还要修改配置文件** : 修改 leanote/conf/app.conf:

```
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username=root # if not exists, please leave blank
db.password=abc123 # if not exists, please leave blank
```

## 为Leanote指定超级管理员帐户(admin用户)
Leanote默认超级管理员为admin, 且一旦不小心修改了username则不能改回. 此时可修改配置文件app.conf, 比如指定用户life为超级管理员, 修改或/添加一行:
```
adminUsername=life
```


## 为Leanote配置https

### 1. 生成SSL证书
可以在网上买一个, 或者自己做一个.
这里有一个shell脚本可以自动生成证书: (cert.sh)
```
#!/bin/sh

# create self-signed server certificate:

read -p "Enter your domain [www.example.com]: " DOMAIN

echo "Create server key..."

openssl genrsa -des3 -out $DOMAIN.key 1024

echo "Create server certificate signing request..."

SUBJECT="/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=$DOMAIN"

openssl req -new -subj $SUBJECT -key $DOMAIN.key -out $DOMAIN.csr

echo "Remove password..."

mv $DOMAIN.key $DOMAIN.origin.key
openssl rsa -in $DOMAIN.origin.key -out $DOMAIN.key

echo "Sign SSL certificate..."

openssl x509 -req -days 3650 -in $DOMAIN.csr -signkey $DOMAIN.key -out $DOMAIN.crt
```

假设得到了两个文件: `a.com.crt`, `a.com.key`

### 2. 配置Nginx
假设Leanote运行的端口是9000, 域名为a.com, 那么nginx.conf可以配置如下:
```
# 本配置只有http部分, 不全
http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
    
    upstream  a.com  {
        server   localhost:9000;
    }

    # http
    server
    {
        listen  80;
        server_name  a.com;
        
        # 强制https
        # 如果不需要, 请注释这一行rewrite
        rewrite ^/(.*) https://jp_linode2.com/$1 permanent;
        
        location / {
            proxy_pass        http://a.com;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
    
    # https
    server
    {
        listen  443 ssl;
        server_name  a.com;
        ssl_certificate     /root/a.com.crt; # 修改路径, 到a.com.crt, 下同
        ssl_certificate_key /root/a.com.key;
        location / {
            proxy_pass        http://a.com;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```


## Import of github.com/revel/revel/modules/testrunner failed

1.0版之前(beta版)会遇到这个问题, 1.0采用revel-0.12, 所以不会遇到这个问题.

```
Failed to load module.  Import of github.com/revel/revel/modules/testrunner failed: cannot find package "github.com/revel/revel/modules/testrunner" in any of:
	/Users/life/app/go1.4/src/github.com/revel/revel/modules/testrunner (from $GOROOT)
	/Users/life/Documents/Go/package_base/src/github.com/revel/revel/modules/testrunner (from $GOPATH)
```
revel 0.12 版配置不一样, 请修改app.conf
```
module.static=github.com/revel/modules/static		 	
module.testrunner=github.com/revel/modules/testrunner	 
```

## 开发版如何更新leanote?

可以使用git pull得到leanote上最新版本, 如果你已修改了leanote, 可以先fetch(推荐使用fetch的方式)最新到本地, 再与本地的合并. 如:
```
git fetch origin master:tmp # 得到远程最新版本, 别名为tmp
git diff tmp # 查看tmp与本地的不同
git merge tmp # 合并到本地
```
如果不能用git方式同步源码, 请下载 https://github.com/leanote/leanote 

1. 请先备份leanote之前的目录, 以防万一
2. 将下载好的替换之前的leanote
3. 将之前版本下的
    * /public/upload/ 目录
    * /files/ 目录
    * /conf/app.conf
移到新版下相应位置.

重启Leanote.

如果运行有问题, 如 "cannot find package "github.com/PuerkitoBio/goquery" in any of:..." 类似的信息, 原因是Leanote增加了新的依赖, 此时可以使用go get命令下载新包, 如下载"github.com/PuerkitoBio/goquery"
```
go get github.com/PuerkitoBio/goquery
```

或下载依赖包与源码全集: https://github.com/leanote/leanote-all 

## 二进制版如何更新leanote?
请下载最新的leanote二进制版, 将之前版本下的

* /public/upload/ 目录
* /files/ 目录
* /conf/app.conf

移到新版下相应位置.

在新版下运行leanote.

## 为什么需要 site.url
site.url是外网可访问的域名, 比如你可以配置为http://a.com, 但在运行leanote可以设端口为9000, 再通过Nginx转发到9000. 如果外网地址是80端口, 请不要填写http://a.com:80, 而只要为http://a.com即可!

site.url用于生成笔记内的图片/附件路径.

若使用nginx转发到https方式部署leanote，site.url需要配置成https://a.com ；否则在博客页面输出的css和js是以http链接形式展现在html中，高版本浏览器比如firefox会直接block掉这部分内容，从而页面显示不正常。更多信息请查看https://github.com/leanote/leanote/issues/228

## 客户端不能同步图片
请确保conf/app.conf的`site.url`和在客户端登录时填写的自建服务地址相同!

## 导出PDF配置 wkhtmltopdf
Leanote的PDF导出使用了wkhtmltopdf, 所以需要先安装wkhtmltopdf, 然后以管理员身份登录Leanote管理后台配置wkhtmltopdf路径.

下文讲解如何安装wkhtmltopdf.

### 源码编译安装
源码编译安装不会出现依赖的问题，wkhtmltopdf提供的编译脚本已经包含了相关依赖库的下载和安装过程，具体编译请查看[wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)的源码包中的install.md文件。

### 二进制版本wkhtmltopdf安装 [推荐]
下载[wkhtmltopdf](http://wkhtmltopdf.org/downloads.html),选择对应的版本，下载下来后通过rpm -ivh命令进行安装。

然后就会看到缺少一些依赖，所以要安装一些依赖库。
#### 安装wkhtmltopdf的依赖包：
```
yum install -y fontconfig libX11 libXext libXrender xorg-x11-fonts-Type1 xorg-x11-fonts-75dpi libpng
```
安装完后再通过rpm -ivh命令安装wkhtmltopdf。

### 测试wkhtmltopdf
使用命令**wkhtmltopdf http://google.com google.pdf**，若没有问题则安装ok。若提示缺少库无法运行，则可以通过**ldd wkhtmltopdf**，查看缺少什么依赖库，然后再安装缺少的依赖库。

可能部分机器会遇到libpng提示缺少的问题，原因是直接安装的版本过高，某些版本wkhtmltopdf预编译版本依赖的libpng版本要低一点，通过yum provides \*libpng12.so.0\*，找到对应的版本，然后在进行安装**yum install libpng12-1.2.50-6.el7.x86_64**。

### 中文问题
在测试了将google页面导出成pdf，并检查pdf内容正常后，可以试一下使用命令**wkhtmltopdf http://baidu.com baidu.pdf**看一下中文页面是否能正常导出。若导出后中文内容为空白则是linux系统缺少中文字体。可以从windwows机器的windwows/fonts目录拷贝几个中文字体文件(微软雅黑之类的)到linux系统的/usr/share/fonts目录，然后再去测试下导出百度页面。


### 不能导出PDF?

如果导出PDF没有出现下载文件的提示, 证明导出出错, 请按照以下步骤检查:

1. wkhtmltopdf安装是否成功? 请通过 `wkhtmltopdf http://baidu.com baidu.pdf` 检查
2. wkhtmltopdf的路径在Leanote是否正确设置? 请复制你设置的路径到控制台导出试试, 比如 `/usr/bin/wkhtmltopdf http://baidu.com baidu.pdf`
3. 导出还有问题? 请检查你的conf/app.conf中设置的`site.url`是否可以访问, Leanote导出PDF时会访问笔记链接, 这个链接的前缀就是 `site.url`

## 不能通过IP访问
2.6版默认绑定localhost, 不能通过ip访问Leanote,

请修改 `app.conf`
```
http.addr=0.0.0.0 # listen on all ip addresses
```
重启Leanote