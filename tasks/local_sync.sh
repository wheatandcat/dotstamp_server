#!/bin/sh
rootDir="/vagrant/Documents/git/dotstamp_server/"
destDir="/home/vagrant/go/src/dotstamp_server/"
targetDirList=("conf/" "static/" "controllers/" "db/" "models/" "resources/" "routers/" "tests/" "utils/" "views" "main.go" "Gomfile" "runner.conf")

for item in ${targetDirList[@]}; do
  rsync -avr --delete $rootDir$item $destDir$item
done
