目前只支持Linux

## 安装依赖
```
$ go get -u -v github.com/gin-gonic/gin
$ go get -u -v github.com/jinzhu/gorm
$ go get -u -v github.com/jinzhu/gorm/dialects/sqlite
$ go get -u -v github.com/gin-contrib/cors
$ go get -u -v github.com/gin-contrib/gzip
$ go get -u -v github.com/lambda-zhang/systemmonitor
$ go get -u -v github.com/gorilla/websocket

$ npm install -g vue-cli
$ cd webpage/
$ npm install
```

## 运行源码
```
$ cd webpage/
$ npm run build
$ cd ..
$ go run main.go
```


## 编译二进制
```
for armd64:
$ CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main.amd64 main.go

for arm64(debug):
$ CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o main.arm64 -v  -ldflags "-linkmode external -extldflags -static" main.go

for arm64(release):
$ CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o main.arm64 -v  -ldflags "-w -s -linkmode external -extldflags -static" main.go

for arm(debug):
$ CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabi-gcc go build -o main.arm -v  -ldflags "-linkmode external -extldflags -static" main.go

for arm(release):
$ CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabi-gcc go build -o main.arm -v  -ldflags "-w -s -linkmode external -extldflags -static" main.go
```


## 运行起来之后在浏览器打开http://127.0.0.1:9000
![截图1](https://github.com/lambda-zhang/systemmonitor-web/blob/master/webpage/static/images/screenshot1.png)
![截图2](https://github.com/lambda-zhang/systemmonitor-web/blob/master/webpage/static/images/screenshot2.png)


## 调试时候检查数据
```
$ sqlite3 systemmonitor.db
sqlite> .database
seq  name             file
---  ---------------  ----------------------------------------------------------
0    main             /data/lambda/systemmonitor-web/systemmonitor.db

sqlite> .tables
os        products

sqlite> .mode column
sqlite> .header on
sqlite> select * from os;
id          up_time     start_time  use_permillage  arch        os          kernel_version  kernel_hostname  num_cpu
----------  ----------  ----------  --------------  ----------  ----------  --------------  ---------------  ----------
1           1           1           1               1           1           1               1                1

sqlite> .quit
```
