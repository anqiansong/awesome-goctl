#/bin/bash
wd=$(pwd)
echo "wd: $wd"


go version
GOPROXY=https://goproxy.cn,direct go install github.com/zeromicro/go-zero/tools/goctl@v1.3.8
goctl --version

