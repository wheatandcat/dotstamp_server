# dotstamp_server

<img src="https://raw.githubusercontent.com/wheatandcat/dotstamp_client/master/dist/images/common/about.png" data-canonical-src="https://raw.githubusercontent.com/wheatandcat/dotstamp_client/master/dist/images/common/about.png" width="200" />

## 概要
.stampのサーバーサイド　  
webサービス：[.stamp](http://dotstamp.com/)

## projectリポジトリ一覧
* サーバーサイド:[dotstamp_server](https://github.com/wheatandcat/dotstamp_server)
* クライアントサイド：[dotstamp_client](https://github.com/wheatandcat/dotstamp_client)
* 環境構築：[dotstamp_ansible](https://github.com/wheatandcat/dotstamp_ansible)
* デプロイスクリプト：[dotstamp_deploy_script](https://github.com/wheatandcat/dotstamp_deploy_script)
* デプロイ環境構築：[dotstamp_deploy_ansible](https://github.com/wheatandcat/dotstamp_deploy_ansible)

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
全体テスト
```
go test -p 1 $(glide novendor)
```
## その他コマンド
指定部分のみテスト
```
go test -p 1 ./models/ -cover TestModel
```
テスト & 監視
```
goconvey -packages 1
```
db確認
```
dbweb -home=$GOPATH/src/github.com/go-xorm/dbweb/
```
## ライセンス
BSDライセンス
