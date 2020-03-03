#检查 goApiFrame 是否存在并是一个文件
if [ ! -f goApiFrame ];then
make
fi
nohup ./goApiFrame conf/conf.ini 2>&1 &
