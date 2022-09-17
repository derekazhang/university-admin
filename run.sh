rm -rf src/logs/*

export GOPROXY=https://goproxy.cn

go mod tidy

GOCMD="GO111MODULE=on go"

eval "${GOCMD} run src/main.go"