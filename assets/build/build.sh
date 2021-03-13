#!/bin/bash
set -e

OS=${1:-"linux"}
ARCH=${2:-"amd64"}
# output path to store leanote output files
OUTPUTPATH=${3:-"./output/leanote"}

# 当前路径
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

##======================
# 1. 目录准备工作
##======================

rm -rf $OUTPUTPATH
mkdir -p $OUTPUTPATH/app
mkdir -p $OUTPUTPATH/conf
mkdir -p $OUTPUTPATH/bin

OUTPUTPATH=$(realpath ${OUTPUTPATH})

##=================================
# 2. build
##=================================


echo build-$OS-$ARCH
if [[ $OS == "linux" || $OS == "darwin" ]]; then
	suffix=""
	if [ $ARCH = "arm" ]
	then
		cp ${SCRIPTPATH}/scripts/run-arm.sh $OUTPUTPATH/bin/run.sh
	else
		cp ${SCRIPTPATH}/scripts/run-$OS-$ARCH.sh $OUTPUTPATH/bin/run.sh
	fi
else
	suffix=".exe"
	cp ${SCRIPTPATH}/scripts/run.bat $OUTPUTPATH/bin/
fi

function generate() {
	GOPATH=$(go env GOPATH)
	if [[ -z "${GOPATH}" ]]; then
		GOPATH="/tmp/gopath"
	fi

	GOSRCPATH="${GOPATH}/src/${1}"
	if [[ ! -e "${GOSRCPATH}" ]]; then
		mkdir -p "${GOSRCPATH}"
	fi

	BUILDSRCPATH=$(realpath "${SCRIPTPATH}/../../")
	if [[ "$(realpath ${BUILDSRCPATH})" != "$(realpath ${GOSRCPATH})" ]]; then
		# rm -rf ${GOSRCPATH}
		# ln -s ${BUILDSRCPATH} ${GOSRCPATH}
		echo  "${BUILDSRCPATH}" != "${GOSRCPATH}"
	fi

	go mod vendor
	GOPATH=${GOPATH} GO111MODULE=off go run ${GOSRCPATH}/assets/build/scripts/generate.go
}

# if [[ ! -e "${SCRIPTPATH}/../../app/tmp" ]]; then 
	generate "github.com/coocn-cn/leanote"
# fi

GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUTPATH/bin/leanote-$OS-$ARCH$suffix" ${SCRIPTPATH}/../../app/tmp

##==================
# 3. 复制
##==================

cd "$SCRIPTPATH"

# bin
cp -r ./assets/src $OUTPUTPATH/bin

# others
cp -r ./assets/public $OUTPUTPATH
cp -r ./assets/messages $OUTPUTPATH
cp -r ./assets/mongodb_backup $OUTPUTPATH

# views
cp -r ../../app/views $OUTPUTPATH/app

# conf
cp ../../conf/routes $OUTPUTPATH/conf/
cp ../../conf/app.conf $OUTPUTPATH/conf/

# 处理app.conf, 还原配置
sed -i 's/db.dbname=leanote.*#/db.dbname=leanote #/' $OUTPUTPATH/conf/app.conf

cd - >/dev/null

# delete some files
rm -f $OUTPUTPATH/public/.DS_Store
rm -f $OUTPUTPATH/public/config.codekit
rm -rf $OUTPUTPATH/public/.codekit-cache
rm -rf $OUTPUTPATH/public/tinymce/classes
rm -rf $OUTPUTPATH/public/upload
mkdir $OUTPUTPATH/public/upload