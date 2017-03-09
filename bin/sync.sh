#!/bin/sh

#cwd=`dirname "${0}"`
#expr "${0}" : "/.*" > /dev/null || cwd=`(cd "${cwd}" && pwd)`

#port="10020"
#user="worker"
#ip="160.16.118.176"


#project=("blue", "green")
#staticDirList=("static/files")

#rootDir="/Users/iinoyouhei/go/src/dotstamp_server/"

#for dir in ${project[@]}; do
#    for item in ${staticDirList[@]}; do
#        target="${user}@${ip}:/project/dotstamp_server/${dir}/${item}"


#        echo "rsync -arv -e 'ssh -p 10020' --exclude='.*' --exclude='*.*' $rootDir$item $target"
#        ${cmd}
#    done
#done

#rsync -arv -e "ssh -p 10020"  --exclude=".*" --exclude="*.*" static/files/ worker@160.16.118.176:/project/dotstamp_server/blue/static/files
#rsync -arv -e 'ssh -p 10020' --exclude='.*' --exclude='*.*' static/files/ worker@160.16.118.176/project/dotstamp_server/blue/static/files
#serverDirList = ("dotstamp_server" "tasks")

#rsync -arv -e "ssh -p 10020" --exclude=".*" --exclude="*.*" /Users/iinoyouhei/go/src/dotstamp_server/static/files worker@160.16.118.176:/project/dotstamp_server/blue/static/files
#rsync -arv -e "ssh -p 10020" --exclude=".*" --exclude="*.*" /Users/iinoyouhei/go/src/dotstamp_server/static/files worker@160.16.118.176:/project/dotstamp_server/blue/static/files
#rsync -arv -e "ssh -p 10020" --exclude=".*" --exclude="*.*" /Users/iinoyouhei/go/src/dotstamp_server/static/files worker@160.16.118.176:/project/dotstamp_server/green/static/files
#rsync -arv -e "ssh -p 10020" --exclude="*.go" /Users/iinoyouhei/go/src/dotstamp_server/dotstamp_server worker@160.16.118.176:/project/dotstamp_server/dotstamp_server
#rsync -arv -e "ssh -p 10020" --exclude="*.go" /Users/iinoyouhei/go/src/dotstamp_server/tasks worker@160.16.118.176:/project/dotstamp_server/tasks
#rsync -arv -e "ssh -p 10020" --exclude="*.go" /Users/iinoyouhei/go/src/dotstamp_server/main worker@160.16.118.176:/project/dotstamp_server/main
rsync -arv -e "ssh -p 10020" --exclude="*.go" /Users/iinoyouhei/go/src/dotstamp_server/conf worker@160.16.118.176:/project/dotstamp_server/conf
