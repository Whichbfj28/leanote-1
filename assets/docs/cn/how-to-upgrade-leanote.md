
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

如果运行有问题
- 如 "cannot find package "github.com/PuerkitoBio/goquery" in any of:..." 类似的信息, 原因是Leanote增加了新的依赖, 此时可以使用go get命令下载新包, 如下载"github.com/PuerkitoBio/goquery"
```
go get github.com/PuerkitoBio/goquery
```
- 如 "cannot find package "golang.org/x/crypto/bcrypt" in any of:" 类似的信息, 原因是Leanote增加了新的依赖, 此时可以使用以下命令下载新包。 如下载"github.com/golang/crypto"
```
cd $GOPATH/src/golang.org/x/
git clone https://github.com/golang/crypto.git
```
或下载依赖包与源码全集: https://github.com/leanote/leanote-all 

## 二进制版如何更新leanote?
请下载最新的leanote二进制版, 将之前版本下的

* /public/upload/ 目录
* /files/ 目录
* /conf/app.conf

移到新版下相应位置.

在新版下运行leanote.
