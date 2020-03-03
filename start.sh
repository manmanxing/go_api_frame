#!/usr/bin/env bash
#if [ $# -ne 1 ];then
#    echo "Paras Error!"
#    echo "Usage: bash ${0} configPath"
#    exit 1
#fi
if [ ! -f goApiFrame ];then
#make clean
make
fi

#echo ${1}
nohup ./goApiFrame conf/conf.ini 2>&1 &
