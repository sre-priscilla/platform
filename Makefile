.PHONY: start
start:
	go run cmd/main.go

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/platform cmd/main.go

.PHONY: image
image:
	cd build && docker build -t registry.cn-shanghai.aliyuncs.com/ayane/platform .