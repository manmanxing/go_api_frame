#!/usr/bin/env bash
#if [ $# -ne 1 ];then
#    echo "Paras Error!"
#    echo "Usage: bash ${0} configPath"
#    exit 1
#fi
#检查 goApiFrame 是否存在并是一个文件
if [ ! -f goApiFrame ];then
#make clean
make
fi

#echo ${1}
nohup ./goApiFrame  2>&1 &
