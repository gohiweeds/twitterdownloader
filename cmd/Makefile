BUILD_NAME:=twitterdownloader
BUILD_VERSION:=1.0
SOURCE:=*.go
LDFLAGS:=-ldflags "-X main.Version=${BUILD_VERSION}"

all: deps build install

deps:
	#安装依赖
	#[ -x glide ] && glide install || yum install glide

test: 
	go test

build: test
	go build -o ${BUILD_NAME} ${SOURCE}

build_linux: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_NAME} ${SOURCE}

build_win: test
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BUILD_NAME} ${SOURCE}

build_macos: test
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BUILD_NAME} ${SOURCE}

install: deps build
	go install
	#生成配置文件等
	#cp app.conf.example /etc/app.conf
#release: build build_win build_macos
	
clean:
	go clean

.PHONY: all deps test build build_linux build_win install clean
