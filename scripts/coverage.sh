#!/bin/sh

controllerDirs=$GOPATH/src/github.com/wheatandcat/dotstamp_server/tests/controllers/*
for filepath in $controllerDirs; do
    if [ -d $filepath ]; then
        dir=`basename ${filepath}`
        echo $dir
        go test -p 1 -coverprofile=./coverage/controllers.$dir.coverage.out  -coverpkg=./controllers/$dir/... ./tests/controllers/$dir/...
    fi
done

list=(models tasks utils)
for item in ${list[@]}; do
    echo $item
    dirs=$GOPATH/src/github.com/wheatandcat/dotstamp_server/$item/*
    go test -p 1 -coverprofile=./coverage/$item.coverage.out./$dir/

    for filepath in $dirs; do
        if [ -d $filepath ]; then
            dir=`basename ${filepath}`
            echo $item/$dir
            go test -p 1 -coverprofile=./coverage/$item.$dir.coverage.out ./$item/$dir/
        fi
    done
done

rm -rf ./coverage/all.coverage.out
rm -rf ./coverage/all.tmp.coverage.out
rm -rf ./coverage/tmp.coverage.out

find ./coverage/ | grep coverage.out | xargs cat | grep -v "mode:" | sort -r >> ./coverage/all.tmp.coverage.out
echo "mode: set" > ./coverage/tmp.coverage.out
cat ./coverage/tmp.coverage.out ./coverage/all.tmp.coverage.out > ./coverage/all.coverage.out
