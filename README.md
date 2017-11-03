# dotstamp_server

[![Build Status](https://travis-ci.org/wheatandcat/dotstamp_server.svg?branch=master)](https://travis-ci.org/wheatandcat/dotstamp_server)
[![DB Document](https://img.shields.io/badge/database-doc-brightgreen.svg)](https://wheatandcat.github.io/dotstamp_server/db/stamp_test.html)
[![API Document](https://img.shields.io/badge/restful%20api-doc-green.svg)](https://wheatandcat.github.io/dotstamp_server/apiary.html#)
[![API Testing](https://img.shields.io/badge/API%20Testing-dredd-green.svg)](https://github.com/apiaryio/dredd)
[![BSD License](https://img.shields.io/badge/license-BSD-blue.svg)](LICENSE)

<img src="https://raw.githubusercontent.com/wheatandcat/dotstamp_client/master/dist/images/common/about.png" data-canonical-src="https://raw.githubusercontent.com/wheatandcat/dotstamp_client/master/dist/images/common/about.png" width="200" />

## 概要
.stampのサーバーサイド　  
webサービス：[.stamp](http://dotstamp.com/)

## projectリポジトリ一覧
* service
  * backend:[dotstamp_server](https://github.com/wheatandcat/dotstamp_server)
  * frontend：[dotstamp_client](https://github.com/wheatandcat/dotstamp_client)
  * ansibl：[dotstamp_ansible](https://github.com/wheatandcat/dotstamp_ansible)
  * deploy_script：[dotstamp_deploy_script](https://github.com/wheatandcat/dotstamp_deploy_script)
  * deploy_ansible：[dotstamp_deploy_ansible](https://github.com/wheatandcat/dotstamp_deploy_ansible)
* admin
  * backend:[dotstamp_server](https://github.com/wheatandcat/dotstamp_admin_server)
  * frontend：[dotstamp_client](https://github.com/wheatandcat/dotstamp_admin_client)

## 必要なもの
* golang
* MariaDB
* Redis
## 環境構築
* [ローカル環境構築](https://github.com/wheatandcat/dotstamp_ansible#ローカル環境構築手順-)
## 実行手順
リポジトリをclone
```
cd $GOPATH/src/
git clone git@github.com:wheatandcat/dotstamp_server.git
cd dotstamp_server
```
ライブラリ導入
```
glide install
```
DBマイグレーション
```
goose up
goose -env test up
```
実行 & 監視
```
bee run
```
バイナリ実行
```
./dotstamp_server
```
アクセス(vagrantから実行した場合)
```
http://192.168.33.10:8080/
```
## その他コマンド
全体テスト
```
go test -p 1 $(glide novendor)
```
指定部分のみテスト
```
go test -p 1 ./models/ -cover TestModel
```
テスト & 監視
```
goconvey -packages 1
```
全体カバレッジ出力 & 出力
```
sh scripts/coverage.sh
go tool cover -func=coverage/all.coverage.out
```
db確認
```
dbweb -home=$GOPATH/src/github.com/go-xorm/dbweb/
```
db定義書生成
```
mysqldump --no-data --xml -uroot stamp_test > db/stamp_test.xml
xsltproc -o db/stamp_test.html db/style.xsl db/stamp_test.xml
```
APIドキュメント生成
```
aglio -i apiary.apib -o apiary.html
```
dredd
```
sh scripts/dredd.sh && ENV_CONF=test dredd
```

## バッチ  
ファイルパス：tasks/

| ファイル名 | 内容 | 備考 |
|:-----------|:------------|:------------|
| contributionSearch       | 投稿の検索テーブルを一括更新する | 新しい検索処理を追加時の再設定用 |
| contributionTotalFollows       | フォロー数の総数を設定する | crontabで15分毎に実行 |
| makeMovie       | 動画ファイルを作成する |  |
| removeContribution       | 不要ファイルを削除する | 1日1回実行 |

## License
MIT
