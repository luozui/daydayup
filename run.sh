#!/bin/bash

DIR=$(cd "$(dirname "$0")"; pwd) 
cd $DIR/gzhu_no_clock_in

while true ; do
        python3 run.py >$DIR/out.log 2>&1
        if [ $? == 0 ] ; then
                echo The command execute OK!
                break;
        fi
done

echo >> $DIR/out.log
date >> $DIR/out.log

# curl --ssl-reqd --url 'smtps://smtp.qq.com:465' --user '329690971:password' --mail-from '329690971@qq.com' --mail-rcpt '329690971@qq.com' --upload-file out.log -v --login-options AUTH=LOGIN
