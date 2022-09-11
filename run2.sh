#!/bin/bash

DIR=$(cd "$(dirname "$0")"; pwd) 
cd $DIR/gzhu_no_clock_in

PRO_NAME=daydayup
NUM=`ps aux | grep ${PRO_NAME} | grep -v grep |wc -l`
#  echo $NUM
#    少于1，重启进程
if [ "${NUM}" -lt "1" ];then
	echo "${PRO_NAME} was killed"
	cd $DIR
	nohup $DIR/${PRO_NAME} >$DIR/server.log 2>&1 &
fi

# curl --ssl-reqd --url 'smtps://smtp.qq.com:465' --user '329690971:password' --mail-from '329690971@qq.com' --mail-rcpt '329690971@qq.com' --upload-file out.log -v --login-options AUTH=LOGIN
