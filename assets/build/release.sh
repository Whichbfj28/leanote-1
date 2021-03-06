#!/bin/bash

# release leanote

# name
NAME=${1:-"release"}
# output path to store leanote output files
OUTPUTPATH=${3:-"./output/leanote"}
# release path to store leanote release files
RELEASEPATH=${3:-"./output"}

# 当前路径
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)


##===========
# 打包
##===========

rm -rf $OUTPUTPATH
mkdir -p $OUTPUTPATH

OUTPUTPATH=$(realpath ${OUTPUTPATH})

# 创建一个$NAME的目录存放之
rm -rf $RELEASEPATH/$NAME
mkdir -p $RELEASEPATH/$NAME

RELEASEPATH=$(realpath ${RELEASEPATH})

# $1 = linux
# $2 = 386, amd64
function tarRelease()
{
	# 编译
	${SCRIPTPATH}/build.sh $1 $2 ${OUTPUTPATH}

	# 打包
	echo tar-$1-$2
	
	if [[ $1 == "linux" || $1 == "darwin" ]]; then
		if [[ $2 == "arm" ]]; then
			cp $SCRIPTPATH/scripts/run-arm.sh $OUTPUTPATH/bin/run.sh
		else
			cp $SCRIPTPATH/scripts/run-$1-$2.sh $OUTPUTPATH/bin/run.sh
		fi
	else
		cp $SCRIPTPATH/scripts/run.bat $OUTPUTPATH/bin/
	fi
	
	tar -czf $RELEASEPATH/$NAME/leanote-$1-$2-$NAME.bin.tar.gz -C "$OUTPUTPATH" .
}

tarRelease "darwin" "amd64";

tarRelease "windows" "386";
tarRelease "windows" "amd64";

tarRelease "linux" "arm";
tarRelease "linux" "386";
tarRelease "linux" "amd64";
