#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

# set link

path="$SCRIPTPATH/src/github.com/coocn-cn"
if [ ! -d "$path" ]; then
	mkdir -p "$path"
fi
rm -rf $SCRIPTPATH/src/github.com/coocn-cn/leanote # 先删除
ln -s ../../../../ $SCRIPTPATH/src/github.com/coocn-cn/leanote

# set GOPATH
export GOPATH=$SCRIPTPATH

script="$SCRIPTPATH/leanote-linux-386"
chmod 777 $script
$script -importPath github.com/coocn-cn/leanote