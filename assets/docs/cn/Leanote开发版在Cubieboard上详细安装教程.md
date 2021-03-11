Cubieboard是由深圳方糖科技公司开发，搭载AllWinner A20芯片，1GB内存，10/100M自适应网络接口。同时，由于Cubieboard支持SATA接口，所以Leanote程序和mongodb可以存放在外接的HDD或SSD硬盘上。

准备工具：

1. Cubieboard2 Dual-Card
2. 一张至少8GB容量的TF卡
3. Cubieboard2的SATA线一根和SATA接口硬盘一块
4. USB转TTL串口线或TTL转RS232串口线

Cubieboard2 Dual-Card这款板子把Nand Flash去掉，换成两个TF卡座，这样Flash的容量就不再受Nand Flash的容量和坏块的限制。两个TF卡座，可以一个用来安装系统，一个用来存储用户数据，但是TF的读写速度不太如意，建议使用SATA外接硬盘。USB转TTL线是以防openssl软件包损坏后，可以通过串口进入系统修复有问题的软件包。

安装步骤：

1. 安装系统
2. 准备运行环境
3. 安装Leanote
4. 运行Leanote

# 安装系统

Cubieboard2 Dual-Card支持多种系统，本次使用的系统镜像是debian-server-cb2-bootcard-hdmi-v1.1.img，可从[cubieboard](cubieboard.org)官网下载到，[下载地址](http://dl.cubieboard.org/model/cubieboard2-dualcard/Image/debian-server-v1.1/)。安装方法，参考此文档[Linux Card installation](http://dl.cubieboard.org/model/cubieboard2-dualcard/Doc/debian-server/Linux-card-installation.pdf)。

# 准备运行环境

Cubieboard2 从TF启动后，从路由器查看到其IP后，使用cubie用户登录SSH连接。首先要更新debian的软件包，在更新之前请先替换新的软件源：

    deb http://mirrors.ustc.edu.cn/debian/ wheezy main non-free contrib
    deb http://mirrors.ustc.edu.cn/debian/ wheezy-updates main non-free contrib
    deb http://mirrors.ustc.edu.cn/debian/ wheezy-backports main non-free contrib
    deb-src http://mirrors.ustc.edu.cn/debian/ wheezy main non-free contrib
    deb-src http://mirrors.ustc.edu.cn/debian/ wheezy-updates main non-free contrib
    deb-src http://mirrors.ustc.edu.cn/debian/ wheezy-backports main non-free contrib
    deb http://mirrors.ustc.edu.cn/debian-security/ wheezy/updates main non-free contrib

Cubieboard使用arm架构的软件包，国内debian软件源支持不多。

    sudo aptitude update
    sudo aptitude upgrade
    sudo aptitude install git mongodb golang


# 安装Leanote

首先下载Leanote开发版的代码，推荐使用Leanote-all的，有些包因为网络问题无法连接，而Leanote-all已包含所有库，可以直接使用。下载地址 [leante-all-master.zip](https://github.com/leanote/leanote-all/archive/master.zip)。解压后，将src目录下载所有移动至gopackage/src下

    mount /dev/sda1 /media/data1
    wget https://github.com/leanote/leanote-all/archive/master.zip
    unzip master.zip
    mkdir -p gopackage/src
    mv leanote-all-master/src/* gopackage/src
    
下面准备mongodb数据库，并把Leanote的数据导入到mongodb。

    cd /media/data1
    mkdir mongodbs
    sudo chown mongodb:mongodb mongodbs
    sudo vi /etc/mongodb.cnf

编辑mongodb.conf文件，修改dbpath，启用port。
    
    dbpath=/media/data1/mongodbs
    port = 27017

然后就可以启动mongodb服务，使用mongo命令查看服务是否启动正常。

    sudo service mongodb start
    mongo
    > version()

如果输入mongo后提示无法连接，请查看mongodb的日志，可找到具体原因。

接下来是导入Leanote的数据库

    mongorestore -h localhost -d leanote --directoryperdb  /media/data1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data
    
执行完成后，可以使用mongo命令，进入leanote数据库，查看是否正确导入

    mongo
    > shwo dbs
    > use leanote
    > show collections

# 运行Leanote

运行Leanote前，需要修改访问Leanote的IP地址和端口

    cd /media/data1
    vi gopackage/src/github.com/leanote/leanote/conf/app.conf
    site.url=http://192.168.2.159:9000
    
然后设置GOPATH后，就可以运行Leanote了

    export GOPATH=/media/data1/gopackage
    export PATH=$GOPATH/bin:$PATH
    revel run github.com/leanote/leanote
    

